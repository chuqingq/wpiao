package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

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
	Key         string // 可以唯一标识一个投票的 TODO
	Url         string // 短URL
	Supervoteid string
	Info        map[string]interface{} // 投票信息。包括活动标题、到期时间、投票对象等
	Item        map[string]interface{} // 投的对象
	Votes       uint64                 // 票数
	Speed       uint64                 // TODO 暂未使用。每分钟的票数
	CurVotes    uint64
}

func GetKeyFromUrl(voteUrl string) string {
	// 解析longUrl中的参数
	u, err := url.Parse(voteUrl)
	if err != nil {
		log.Printf("getKeyFromUrl: parse url error: %v", err)
		return ""
	}

	values := u.Query()
	if values.Get("__biz") == "" {
		log.Printf("getKeyFromUrl: __biz is empty")
		return ""
	}

	key := "__biz=" + values.Get("__biz") + "&mid=" + values.Get("mid") + "&idx=" + values.Get("idx") + "&sn=" + values.Get("sn")
	return key
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
