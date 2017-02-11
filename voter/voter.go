package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const SHORT_URL = "http://mp.weixin.qq.com/s/WEBkpBjBdOAIXxu9fknV9w"
const ITEM = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`
const DST_VOTES = 1
const VOTE_URL = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&key=f6fc65d37e8c2007e879f47762586e65a02d8fbd5b84db235e00e511b8101f887e892a2554674628ca531decec74f300247b10a9d1bddcb0db5ed37662159345e43c794bdb7046a6a6c53cd203b232d1&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=EnayxJ3mRIUH%2BQl8MDq4Bjq1qQJiB0M4Od8lSTPh3ejMZ1VSt03lQLCWB0LI5dKT"

var gVoteInfos = VoteInfos{}

func main() {
	http.HandleFunc("/api/parseurl", ParseUrl)
	http.HandleFunc("/api/submittask", SubmitTask)
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

type VoteInfos map[string]*VoteInfo

func (vis VoteInfos) Get(key string) *VoteInfo {
	return vis[key]
}

func (vis VoteInfos) Set(key string, vi *VoteInfo) {
	vis[key] = vi
}

func (vis VoteInfos) Del(key string) {
	delete(vis, key)
}

type VoteInfo struct {
	Url         string
	Key         string // 可以唯一标识一个投票的
	Supervoteid string
	Info        map[string]interface{} // title等信息都在这里
	Item        map[string]interface{} // 投的对象 TODO 改为DstItem
	DstVotes    uint64                 // 目标票数
	CurVotes    uint64                 // 当前票数
}

func NewVoteInfo(shortOrLongUrl string) (*VoteInfo, error) {
	log.Printf("NewVoteInfo shortOrLongUrl: %v", shortOrLongUrl)

	// 设置cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("cookiejar.New() error: %v", err)
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	vi := &VoteInfo{
		Url: strings.Replace(shortOrLongUrl, "https:", "http:", 1),
	}

	resp, err := client.Get(vi.Url)
	if err != nil {
		log.Printf("get shorturl error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %v", err)
		return nil, err
	}

	// 可能是短连接，这时需要拿到长连接（不需请求，直接parse拿到参数即可）
	if strings.Contains(shortOrLongUrl, "/s/") {
		vi.Url = string(getByBound(resBody, []byte(`var msg_link = "`), []byte(`";`)))
		vi.Url = strings.Replace(vi.Url, "https:", "http:", 1)
		vi.Url = strings.Replace(vi.Url, `\x26amp;`, `&`, -1)
		log.Printf("longurl: %v", vi.Url)
		if vi.Url == "" {
			log.Printf("get longurl error")
			return nil, errors.New("get longurl error")
		}
	}

	vi.Supervoteid = string(getByBound(resBody, []byte(`supervoteid=`), []byte(`&`)))
	log.Printf("supervoteid: %v", vi.Supervoteid)
	if vi.Supervoteid == "" {
		log.Printf("supervoteid is empty. maybe url is invalid")
		return nil, errors.New("supervoteid is empty. maybe url is invalid")
	}

	// 解析longUrl中的参数
	u, err := url.Parse(vi.Url)
	if err != nil {
		log.Printf("parse url error: %v", err)
		return nil, err
	}

	values := u.Query()
	vi.Key = "__biz=" + values.Get("__biz") + "&mid=" + values.Get("mid") + "&idx=" + values.Get("idx") + "&sn=" + values.Get("sn")

	// 获取投票信息
	values.Set("supervoteid", vi.Supervoteid)
	values.Set("action", "show")
	showUrl := getNewappmsgvoteShowUrl(values)
	log.Printf("showUrl: %v", showUrl)
	resp, err = client.Get(showUrl)
	if err != nil {
		log.Printf("getNewappmsgvoteShowUrl error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body 2 error: %v", err)
		return nil, err
	}

	voteInfoStr := string(getByBound(resBody, []byte(`var voteInfo=`), []byte(`;`)))
	log.Printf("voteInfoStr: %v", voteInfoStr) // TODO
	vi.Info = make(map[string]interface{})
	err = json.Unmarshal([]byte(voteInfoStr), &vi.Info)
	log.Printf("info: %v", vi.Info)
	if err != nil {
		log.Printf("json.Unmarshal voteInfo error: %v", err)
		return nil, err
	}

	return vi, nil
}

func (vi *VoteInfo) NewVoter(voteUrl string) (*Voter, error) {
	log.Printf("NewVoter voteUrl: %v", voteUrl)
	voteUrl = strings.Replace(voteUrl, "https:", "http:", 1)

	// 设置cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("cookiejar.New() error: %v", err)
		return nil, err
	}

	// 解析其他参数
	u, err := url.Parse(voteUrl)
	if err != nil {
		log.Printf("parse url error: %v", err)
		return nil, err
	}

	return &Voter{
		url: voteUrl,
		// item: item, // TODO 应该放在VoteInfo中
		client: &http.Client{
			Jar: jar,
		},
		values: u.Query(),
		Info:   vi,
	}, nil
}

type Voter struct {
	url    string
	client *http.Client
	values url.Values
	Info   *VoteInfo
}

func (v *Voter) Vote() error {
	err := v.s()
	if err != nil {
		log.Printf("s error: %v", err)
		return err
	}

	err = v.newappmsgvoteShow()
	if err != nil {
		log.Printf("newappmsgvoteShow error: %v", err)
		return err
	}

	err = v.newappmsgvoteVote()
	if err != nil {
		log.Printf("newappmsgvoteVote error: %v", err)
		return err
	}

	// 增加voteInfo的计数
	v.Info.CurVotes += 1
	return nil
}

func (v *Voter) s() error {
	res, err := v.client.Get(v.url)
	if err != nil {
		log.Printf("client.Get(url) error: %v", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("read body error: %v", err)
		return err
	}

	v.values.Set("wxtoken", string(getByBound(resBody, []byte("window.wxtoken = \""), []byte("\";"))))
	log.Printf("wxtoken: %v", v.values.Get("wxtoken"))
	if v.values.Get("wxtoken") == "" {
		log.Printf("get wxtoken error. maybe voteUrl expired")
		return errors.New("get wxtoken error. maybe voteUrl expired")
	}

	return nil
}

func (v *Voter) newappmsgvoteShow() error {
	url := getNewappmsgvoteShowUrl(v.values)
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("newappmsgvoteShow get error: %v", err)
		return err
	}
	defer res.Body.Close()

	// 这里不需要做什么，supervoteid是之前就有的

	return nil
}

func getNewappmsgvoteShowUrl(values url.Values) string { // TODO 合并
	return "http://mp.weixin.qq.com/mp/newappmsgvote?" + values.Encode()
}

func (v *Voter) newappmsgvoteVote() error {
	// TODO 投票对象如何确定？
	v.values.Set("action", "vote")
	v.values.Set("f", "json")
	// v.values.Set("json", v.item)
	item, err := json.Marshal(v.Info.Item)
	if err != nil {
		log.Printf("json.Marshal item error: %v", err)
		return err
	}
	v.values.Set("json", string(item))
	log.Printf("vote values: %v", v.values)
	log.Printf("newappmsgvoteVote formdata: %v", v.values.Encode())

	res, err := v.client.PostForm("https://mp.weixin.qq.com/mp/newappmsgvote", v.values)
	if err != nil {
		log.Printf("newappmsgvoteVote error: %v", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("read resBody error: %v", err)
		return err
	}
	log.Printf("resBody: %v", string(resBody))

	return nil
}

func getByBound(b, left, right []byte) []byte {
	lindex := bytes.Index(b, left)
	if lindex < 0 {
		return nil
	}

	lindex += len(left)
	rindex := bytes.Index(b[lindex:], right)
	if rindex < 0 {
		return nil
	}

	return b[lindex:(lindex + rindex)]
}
