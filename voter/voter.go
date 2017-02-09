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
const VOTE_URL = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&key=f6fc65d37e8c200728d1961ac018cf11978985a26153db4e52999fff8c0752d6acea32e2d784dda5df2b23afba6fca6173dfd974bb08f73dc9b30906521d1a2eb9aa66865b4e034af663c6b4bc8b5bcc&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=EnayxJ3mRIUH%2BQl8MDq4Bjq1qQJiB0M4Od8lSTPh3ejMZ1VSt03lQLCWB0LI5dKT"

func main() {
	// 根据短url来获取到投票信息
	voteInfo, err := NewVoteInfo(SHORT_URL)
	if err != nil {
		log.Fatalf("NewVoteInfo error: %v", err)
	}
	log.Printf("VoteInfo: %+v", voteInfo)
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
	// TODO 按照__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6格式唯一确定一个voteInfo
}

func (vis VoteInfos) Del(key string) {
	delete(vis, key)
}

type VoteInfo struct {
	Url      string
	Key      string                 // 可以唯一标识一个投票的
	Item     map[string]interface{} // 投的对象
	DstVotes uint64                 // 目标票数
	CurVotes uint64                 // 当前票数
}

func NewVoteInfo(shortOrLongUrl string) (*VoteInfo, error) {
	log.Printf("NewVoteInfo shortOrLongUrl: %v", shortOrLongUrl)
	// 先换成http的 TODO

	// 可能是短连接，也可能是长连接 TODO
	longUrl := shortOrLongUrl
	if strings.Contains(shortOrLongUrl, "/s/") {
		resp, err := http.Get(shortOrLongUrl)
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

		longUrl = string(getByBound(resBody, []byte(`var msg_link = "`), []byte(`";`)))
		log.Printf("longurl: %v", longUrl)
		if longUrl == "" {
			log.Printf("get longurl error")
			return nil, errors.New("get longurl error")
		}
	}

	// 解析longUrl中的参数
	u, err := url.Parse(longUrl)
	if err != nil {
		log.Printf("parse longurl error: %v", err)
		return nil, err
	}

	values := u.Query()
	key := "__biz=" + values.Get("__biz") + "&mid=" + values.Get("mid") + "&idx=" + values.Get("idx") + "&sn=" + values.Get("sn")

	return &VoteInfo{
		Url: longUrl,
		Key: key,
	}, nil
}

func (vi *VoteInfo) NewVoter(voteUrl string) (*Voter, error) {
	log.Printf("NewVoter voteUrl: %v", voteUrl)

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
	url    string // 长地址
	client *http.Client
	values url.Values
	Info   *VoteInfo
}

// func NewVoter(surl, item string) (*Voter, error) {
// 	// 设置cookiejar
// 	jar, err := cookiejar.New(nil)
// 	if err != nil {
// 		log.Printf("cookiejar.New() error: %v", err)
// 		return nil, err
// 	}

// 	// 解析其他参数
// 	u, err := url.Parse(surl)
// 	if err != nil {
// 		log.Printf("parse url error: %v", err)
// 		return nil, err
// 	}

// 	return &Voter{
// 		url:  surl,
// 		item: item,
// 		client: &http.Client{
// 			Jar: jar,
// 		},
// 		values: u.Query(),
// 	}, nil
// }

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
		log.Printf("wxtoken error, resBody: %v", string(resBody))
		return errors.New("get wxtoken error")
	}

	return nil
}

func (v *Voter) newappmsgvoteShow() error {
	url := v.getNewappmsgvoteShowUrl()
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("newappmsgvoteShow get error: %v", err)
		return err
	}
	defer res.Body.Close()

	// TODO 看需要取什么内容 supervoteid是之前就有的

	return nil
}

func (v *Voter) getNewappmsgvoteShowUrl() string {
	v.values.Set("action", "show")
	log.Printf("show values: %v", v.values.Encode())

	log.Printf("newappmsgvoteShowUrl: %v", "https://mp.weixin.qq.com/mp/newappmsgvote?"+v.values.Encode())
	return "https://mp.weixin.qq.com/mp/newappmsgvote?" + v.values.Encode()
}

func (v *Voter) newappmsgvoteVote() error {
	// TODO
	// super_vote_item和super_vote_id是提前组装好的，因此只要encodeURIComponent即可
	// formdata = "action=vote&__biz=MzA5NjYwOTg0Nw%3D%3D&uin=MTMwMzUxMjg3Mw%3D%3D&key=82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&f=json&json=%7B%22super_vote_item%22%3A%5B%7B%22vote_id%22%3A684888407%2C%22item_idx_list%22%3A%7B%22item_idx%22%3A%5B%220%22%5D%7D%7D%5D%2C%22super_vote_id%22%3A684888406%7D&idx=1&mid=2650886522&wxtoken=543112670"
	// values := url.Values{}
	// values.Set("action", "vote")
	// values.Set("__biz", v.__biz)
	// values.Set("uin", v.uin) // TODO 是否需要转译？
	// values.Set("key", v.key)
	// values.Set("pass_ticket", v.pass_ticket)
	// values.Set("f", "json")
	// values.Set("json", item)
	// values.Set("idx", v.idx)
	// values.Set("mid", v.mid)
	// values.Set("wxtoken", v.wxtoken)
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
