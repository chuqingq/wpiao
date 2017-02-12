package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const SHORT_URL = "http://mp.weixin.qq.com/s/WEBkpBjBdOAIXxu9fknV9w"
const ITEM = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`
const DST_VOTES = 1
const VOTE_URL = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&key=f6fc65d37e8c2007e879f47762586e65a02d8fbd5b84db235e00e511b8101f887e892a2554674628ca531decec74f300247b10a9d1bddcb0db5ed37662159345e43c794bdb7046a6a6c53cd203b232d1&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=EnayxJ3mRIUH%2BQl8MDq4Bjq1qQJiB0M4Od8lSTPh3ejMZ1VSt03lQLCWB0LI5dKT"

var gVoteInfos = VoteInfos{}

func main() {
	http.HandleFunc("/api/parseurl", ParseUrl)
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
	log.Printf("voteInfo: %+v", voteInfo)

	infoBytes, err := json.Marshal(voteInfo.Info)
	if err != nil {
		log.Printf("marshal info error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	gVoteInfos.Set(voteInfo.Key, voteInfo)
	w.Write(infoBytes)
}

func SubmitTask(w http.ResponseWriter, r *http.Request) {
	// TODO
	log.Printf("/submittask")

	voteResult := r.FormValue("task")
	if voteResult == "" {
		log.Printf("task is empty")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("voteResult: %v", voteResult)
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
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

	log.Printf("remoteaddr: %v", ws.RemoteAddr())
	gWsConns[ws.RemoteAddr().String()] = ws

	// msgtype, msg, err := ws.ReadMessage()
	// if err != nil {
	// 	log.Printf("ws.ReadMessage error: %v", err)
	// 	return
	// }
	// log.Printf("type: %v, content: %v", msgtype, string(msg))

}

func WsWeb(w http.ResponseWriter, r *http.Request) {
	log.Printf("/api/ws/web")
	// TODO
}

// 测试main函数
func main2() {
	// 根据短url来获取到投票信息
	voteInfo, err := NewVoteInfo(VOTE_URL)
	if err != nil {
		log.Fatalf("NewVoteInfo error: %v", err)
	}
	log.Fatalf("VoteInfo: %+v", voteInfo)
	// 前端可以通过voteInfo展示信息，例如标题、活动日期、当前票数等

	// 添加到voteInfos中
	voteInfos := VoteInfos{}
	voteInfos.Set(voteInfo.Key, voteInfo)

	// 前端确定投票对象（可以根据ID）
	item := make(map[string]interface{})
	err = json.Unmarshal([]byte(ITEM), &item)
	if err != nil {
		log.Fatalf("json.Unmarshal ITEM error: %v", err)
	}

	// 根据key（前端有）找到voteInfo，设置item
	key := voteInfo.Key
	voteInfo2 := voteInfos.Get(key)
	voteInfo2.Item = item
	voteInfo2.DstVotes = DST_VOTES

	// 根据账号的url和item来执行
	voteInfo3 := voteInfos.Get(key)
	voter, err := voteInfo3.NewVoter(VOTE_URL)
	if err != nil {
		log.Printf("newvoter error: %v", err)
		return
	}
	log.Printf("Voter: %+v", voter)

	err = voter.Vote()
	log.Printf("vote: %v", err)
}

var gWsConns = map[string]*websocket.Conn{}
