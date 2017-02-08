package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const surl = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&from=singlemessage&isappinstalled=1&key=dd09eb91646892ca1914090e28a382a26ee75d0b3d52ce6c76a835c5190381212c0c6e6085fab6c12b84066d167100a03e328b9de2489c4f19bfc934e1fa305d7d705b8aa49d7a1f3c21d3d3b2dd5250&ascene=1&uin=MzcyMDM3MzU1&devicetype=Windows+7&version=61000603&pass_ticket=x54VPZwfTTyIFzz7u%2BDId8SELfmTs55SaXxe4coMJKEXcjpPtk8d3cx4AvRSpIPA"
const item = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`

func main() {
	voter, err := NewVoter(surl, item)
	if err != nil {
		log.Printf("newvoter error: %v", err)
		return
	}
	log.Printf("voter: %+v", voter)

	err = voter.Vote()
	log.Printf("vote: %v", err)
}

type Voter struct {
	url    string // 长地址
	item   string // 投票内容
	client *http.Client
	values url.Values
	// __biz       string
	// mid         string
	// idx         string
	// sn          string
	// key         string
	// uin         string
	// pass_ticket string
	// supervoteid string
	// wxtoken     string
}

func NewVoter(surl, item string) (*Voter, error) {
	// 设置cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("cookiejar.New() error: %v", err)
		return nil, err
	}

	// 解析其他参数
	u, err := url.Parse(surl)
	if err != nil {
		log.Printf("parse url error: %v", err)
		return nil, err
	}

	return &Voter{
		url:  surl,
		item: item,
		client: &http.Client{
			Jar: jar,
		},
		values: u.Query(),
	}, nil
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
	v.values.Set("json", v.item)
	log.Printf("json: %v", url.QueryEscape(v.item))
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
