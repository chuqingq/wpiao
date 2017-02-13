package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

const SHORT_URL = "http://mp.weixin.qq.com/s/WEBkpBjBdOAIXxu9fknV9w"
const ITEM = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`
const DST_VOTES = 1
const VOTE_URL = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&key=f6fc65d37e8c2007e879f47762586e65a02d8fbd5b84db235e00e511b8101f887e892a2554674628ca531decec74f300247b10a9d1bddcb0db5ed37662159345e43c794bdb7046a6a6c53cd203b232d1&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=EnayxJ3mRIUH%2BQl8MDq4Bjq1qQJiB0M4Od8lSTPh3ejMZ1VSt03lQLCWB0LI5dKT"

var gVoteInfosPrepare = VoteInfos{} // 经过parseurl但未submittask的
var gVoteInfos = VoteInfos{}        // 经过submittask的
var gVoteInfosFinish = VoteInfos{}  // 已经投票完成的 TODO 暂未使用

func main() {
	http.HandleFunc("/api/tasks", Tasks)
	http.HandleFunc("/api/parseurl", ParseUrl)
	http.HandleFunc("/api/submititem", SubmitItem)
	http.HandleFunc("/api/submittask", SubmitTask)

	// TODO websocket1: /api/ws/pc PC端连接，下发任务
	http.HandleFunc("/api/ws/pc", WsPC)
	// TODO websocket2: /api/ws/web web端连接，实时查询状态
	http.HandleFunc("/api/ws/web", WsWeb)
	http.Handle("/", http.FileServer(http.Dir("../web")))

	const addr = ":8080"
	log.Printf("listen at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	log.Printf("Tasks: %+v", gVoteInfos)

	tasks := []map[string]interface{}{}
	for _, info := range gVoteInfos {
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
	log.Printf("/parseurl")

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
	// voteInfo["key"] = voteInfo.Key // 添加一个key作为后续操作的标识
	log.Printf("voteInfo: %+v", voteInfo)

	infoBytes, err := json.Marshal(voteInfo.Info)
	if err != nil {
		log.Printf("marshal info error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	gVoteInfosPrepare.Set(voteInfo.Key, voteInfo)

	w.Write(infoBytes)
}

// 提交投票对象
func SubmitItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("SubmitItem:")

	itemStr := r.FormValue("item")
	if itemStr == "" {
		log.Printf("item is empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	item := map[string]interface{}{}
	err := json.Unmarshal([]byte(itemStr), &item)
	if err != nil {
		log.Printf("json.Unmarshal item error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("item: %v", item)

	// 根据super_vote_id查找voteInfo
	supervoteid := strconv.FormatUint(uint64(item["super_vote_id"].(float64)), 10)
	log.Printf("supervoteid: %v", supervoteid)
	var voteInfo *VoteInfo
	for _, info := range gVoteInfosPrepare {
		if info.Supervoteid == supervoteid {
			voteInfo = info
			break
		}
	}

	if voteInfo == nil {
		log.Printf("voteInfo not found for super_vote_id: %v", supervoteid)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	voteInfo.Item = item
	w.Write([]byte("{}"))
}

// 提交任务
func SubmitTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("SubmitTask:")

	superVoteId := r.FormValue("super_vote_id")
	if superVoteId == "" {
		log.Printf("super_vote_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskStr := r.FormValue("task")
	if taskStr == "" {
		log.Printf("task is empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task := map[string]interface{}{}
	err := json.Unmarshal([]byte(taskStr), &task)
	if err != nil {
		log.Printf("json.Unmarshal task error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("task: %v", task)

	// 根据super_vote_id查找voteInfo
	var voteInfo *VoteInfo
	for _, info := range gVoteInfosPrepare {
		log.Printf("info.Supervoteid: %v", info.Supervoteid)
		if info.Supervoteid == superVoteId {
			voteInfo = info
			break
		}
	}

	if voteInfo == nil {
		log.Printf("voteInfo not found for super_vote_id: %v", superVoteId)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	voteInfo.Votes = uint64(task["votes"].(float64))
	voteInfo.Speed = uint64(task["votespermin"].(float64))
	w.Write([]byte("{}"))

	// 需要把voteinfo从prepare放在voting中
	delete(gVoteInfosPrepare, voteInfo.Key)
	gVoteInfos.Set(voteInfo.Key, voteInfo)

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

func WsWeb(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/ws/web")
	// TODO
}

// // 测试main函数
// func main2() {
// 	// 根据短url来获取到投票信息
// 	voteInfo, err := NewVoteInfo(VOTE_URL)
// 	if err != nil {
// 		log.Fatalf("NewVoteInfo error: %v", err)
// 	}
// 	log.Fatalf("VoteInfo: %+v", voteInfo)
// 	// 前端可以通过voteInfo展示信息，例如标题、活动日期、当前票数等

// 	// 添加到voteInfos中
// 	voteInfos := VoteInfos{}
// 	voteInfos.Set(voteInfo.Key, voteInfo)

// 	// 前端确定投票对象（可以根据ID）
// 	item := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(ITEM), &item)
// 	if err != nil {
// 		log.Fatalf("json.Unmarshal ITEM error: %v", err)
// 	}

// 	// 根据key（前端有）找到voteInfo，设置item
// 	key := voteInfo.Key
// 	voteInfo2 := voteInfos.Get(key)
// 	voteInfo2.Item = item
// 	voteInfo2.Votes = DST_VOTES

// 	// 根据账号的url和item来执行
// 	voteInfo3 := voteInfos.Get(key)
// 	voter, err := voteInfo3.NewVoter(VOTE_URL)
// 	if err != nil {
// 		log.Printf("newvoter error: %v", err)
// 		return
// 	}
// 	log.Printf("Voter: %+v", voter)

// 	err = voter.Vote()
// 	log.Printf("vote: %v", err)
// }

var gWsConns = map[string]*websocket.Conn{}
