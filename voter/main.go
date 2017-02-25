package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "strconv"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

const SHORT_URL = "http://mp.weixin.qq.com/s/WEBkpBjBdOAIXxu9fknV9w"
const ITEM = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`
const DST_VOTES = 1
const VOTE_URL = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&key=f6fc65d37e8c2007e879f47762586e65a02d8fbd5b84db235e00e511b8101f887e892a2554674628ca531decec74f300247b10a9d1bddcb0db5ed37662159345e43c794bdb7046a6a6c53cd203b232d1&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=EnayxJ3mRIUH%2BQl8MDq4Bjq1qQJiB0M4Od8lSTPh3ejMZ1VSt03lQLCWB0LI5dKT"

var gWsConns = map[string]*websocket.Conn{}

func main() {
	// mongo
	err := InitMongo("127.0.0.1")
	if err != nil {
		log.Fatalf("init mongo error: %v", err)
	}

	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionProvider 默认是 memory，目前支持还有 file、mysql、redis 等
	// beego.BConfig.WebConfig.Session.SessionGCMaxLifetime 默认3600秒
	http.HandleFunc("/api/login", Login)
	http.HandleFunc("/api/tasks", Tasks)
	http.HandleFunc("/api/parseurl", ParseUrl)
	http.HandleFunc("/api/submititem", SubmitItem)
	http.HandleFunc("/api/submittask", SubmitTask)

	http.HandleFunc("/api/users", UsersHandle)
	http.HandleFunc("/api/newuser", NewUser)

	// TODO websocket1: /api/ws/pc PC端连接，下发任务
	http.HandleFunc("/api/ws/pc", WsPC)
	http.HandleFunc("/api/vote", PCVote)

	// TODO websocket2: /api/ws/web web端连接，实时查询状态
	http.HandleFunc("/api/ws/web", WsWeb)
	http.Handle("/", http.FileServer(http.Dir("../web")))

	const addr = ":8080"
	log.Printf("listen at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login:")

	user := UserLogin(w, r)
	if user == nil {
		// 如果返回nil，说明失败，内部会回复响应，所以这里直接return
		return
	}

	w.Write([]byte(`{"ret":0,"msg":"login success"}`))
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/tasks:")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	// 获取任务时需要按照user获取
	voteInfos, err := QueryVoteInfosByUser(user.UserName)
	if err != nil {
		log.Printf("QueryVoteInfosByUser error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tasks := []map[string]interface{}{}
	for _, info := range voteInfos {
		task := map[string]interface{}{}
		task["title"] = info.Info["title"]
		task["votes"] = info.Votes
		task["curvotes"] = info.CurVotes
		tasks = append(tasks, task)
	}
	log.Printf("tasks: %+v", tasks)

	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(tasksBytes)
}

// 根据url解析出投票信息
func ParseUrl(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/parseurl:")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	voteUrl := r.FormValue("url")
	if voteUrl == "" {
		log.Printf("url is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("voteUrl: %v", voteUrl)

	// 根据短url来获取投票信息
	voteInfo, err := NewVoteInfo(voteUrl)
	if err != nil {
		log.Printf("NewVoteInfo error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("voteInfo: %+v", voteInfo)

	infoBytes, err := json.Marshal(voteInfo.Info)
	if err != nil {
		log.Printf("marshal info error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(infoBytes)
}

// 提交投票对象
func SubmitItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/submititem:")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	w.Write([]byte("{}"))
}

// 提交任务
func SubmitTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("SubmitTask:")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	// url
	voteUrl := r.FormValue("url")
	if voteUrl == "" {
		log.Printf("url is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// info
	infoStr := r.FormValue("info")
	if infoStr == "" {
		log.Printf("info is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	info := map[string]interface{}{}
	// err := json.Unmarshal([]byte(infoStr), &info)
	err := jsonUnmarshal([]byte(infoStr), &info)
	if err != nil {
		log.Printf("json.Unmarshal info error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("info: %v", info)

	// 取出key
	key := info["key"].(string)
	log.Printf("key: %v", key)
	delete(info, "key")

	// item
	itemStr := r.FormValue("item")
	if itemStr == "" {
		log.Printf("item is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// item := map[string]interface{}{}
	// // err = json.Unmarshal([]byte(itemStr), &item)
	// err = jsonUnmarshal([]byte(itemStr), &item)
	// if err != nil {
	// 	log.Printf("json.Unmarshal item error: %v", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// log.Printf("item: %v", item)
	log.Printf("itemStr: %v", itemStr)

	// task
	taskStr := r.FormValue("task")
	if taskStr == "" {
		log.Printf("task is empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	task := map[string]interface{}{}
	// err = json.Unmarshal([]byte(taskStr), &task)
	err = jsonUnmarshal([]byte(taskStr), &task)
	if err != nil {
		log.Printf("json.Unmarshal task error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("task: %v", task)

	votes, _ := task["votes"].(json.Number).Int64()
	speed, _ := task["votespermin"].(json.Number).Int64()

	voteInfo := &VoteInfo{
		Id: bson.NewObjectId(),
		Url: voteUrl,
		// Key:    GetKeyFromUrl(voteUrl),
		Key:    key,
		Info:   info,
		// Item:   item,
		Item:   itemStr,
		User:   user.UserName,
		Votes:  uint64(votes),
		Speed:  uint64(speed),
		Status: "doing",
	}

	// 写到数据库中
	err = voteInfo.Insert()
	if err != nil {
		log.Printf("voteinfo.insert error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("{}"))

	// 处理任务
	pcCount := len(gWsConns)
	if pcCount == 0 {
		log.Printf("ERROR executer not found")
		return
	}

	for _, wsConn := range gWsConns {
		req := map[string]interface{}{}
		req["cmd"] = "vote"
		req["url"] = voteInfo.Url
		req["votes"] = voteInfo.Votes
		err := wsConn.WriteJSON(req)
		if err != nil {
			log.Printf("ws.WriteJSON error: %v", err)
		}
		log.Printf("dispatch task(%v,%v) to executer(%v)", voteInfo.Url, voteInfo.Votes, wsConn.RemoteAddr().String())
		break // TODO 暂时直接把所有票数都发到第一个pc上去
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsPC(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/ws/pc")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	addr := ws.RemoteAddr().String()
	log.Printf("remoteaddr: %v", addr)
	gWsConns[addr] = ws
	defer delete(gWsConns, addr)

	for {
		msgtype, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("ws.ReadMessage error: %v", err)
			return
		}
		log.Printf("type: %v, content: %v", msgtype, string(msg))
	}
}

// browser把url发过来
func PCVote(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/vote:")

	voteUrl := r.FormValue("url")
	if voteUrl == "" {
		log.Printf("url is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("voteUrl: %v", voteUrl)

	key := GetKeyFromUrl(voteUrl)
	if key == "" {
		log.Printf("get empty key from url %v", voteUrl)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 先回响应
	w.WriteHeader(http.StatusOK)

	voteInfo, err := QueryVoteInfoByKey(key)
	if err != nil {
		log.Printf("QueryVoteInfoByKey(%s) error: %v", key, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	voter, err := voteInfo.NewVoter(voteUrl)
	if err != nil {
		log.Printf("newvoter error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// log.Printf("type of item: %T", voteInfo.Item["super_vote_id"])

	err = voter.Vote()
	if err != nil {
		log.Printf("vote error: %v", err)
		// 如果投票失败，则票数-1
		voteInfo.DecrVotes()
	}

	return
}

func WsWeb(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/ws/web")
	// TODO
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/users")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	users, err := user.QueryAllUsers()
	if err != nil {
		log.Printf("QueryAllUsers error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	by, err := json.Marshal(users)
	if err != nil {
		log.Printf("marshal users(%v) error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(by)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/newuser")

	user := UserLogin(w, r)
	if user == nil {
		return
	}

	if !user.IsAdmin {
		log.Printf("current user is not admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		log.Printf("username is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")
	if password == "" {
		log.Printf("password is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newuser := &User{
		UserName: username,
		Password: password,
		IsAdmin:  false,
	}

	err := newuser.Insert()
	if err != nil {
		log.Printf("user insert error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}
