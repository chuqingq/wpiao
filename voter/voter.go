package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

const url = "https://mp.weixin.qq.com/s?__biz=MzA5NjYwOTg0Nw==&mid=2650886522&idx=1&sn=317f363e12cd7c45e6bbc0de9916a6c6&from=singlemessage&isappinstalled=1&key=82438a29ddf260107913e57d985e8b6d1de73abc6f2481d34844c234e7e5e917df497bfb3b93ab1867c529fb21beecd2fdbbeb7b16c1820304b159f52424bd3530cb1dbac87098aff065d24729104f85&ascene=1&uin=MTMwMzUxMjg3Mw%3D%3D&devicetype=Windows+7&version=61000603&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%2BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%2BN8W"
const item = `{"super_vote_item":[{"vote_id":684888407,"item_idx_list":{"item_idx":["0"]}}],"super_vote_id":684888406}`

func main() {
	vote(url, item)
}

type Voter struct {
	url  string // 长地址
	item string // 投票内容

	client      *http.Client
	supervoteid string // ?
	biz         string
	mid         string
	idx         string
	sn          string
	key         string
	uin         string
	passticket  string
	wxtoken     string
}

func NewVoter(url, item string) *Voter {
	// 设置cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("cookiejar.New() error: %v", err)
		return nil
	}

	// TODO 解析其他参数
	return &Voter{
		client: &http.Client{
			Jar: jar,
		},
		url: url,
	}
}

func (v *Voter) Vote() error {
	err := v.s()
	if err != nil {
		log.Printf("s error: %v", err)
		return err
	}

	// newappmsgvoteShow()
	err = v.newappmsgvoteShow()
	if err != nil {
		log.Printf("newappmsgvoteShow error: %v", err)
		return err
	}

	// newappmsgvoteVote()
	err = v.newappmsgvoteVote()
	if err != nil {
		log.Printf("newappmsgvoteVote error: %v", err)
		return err
	}
}

func (v *Voter) s() error {
	// s()
	// 从url中取出参数:
	// 从响应中取出supervoteid和wxtoken
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("client.Get(url) error: %v", err)
		return err
	}
	defer res.Body.close()

	resBody, err := ioutil.ReadAll()
	if err != nil {
		log.Printf("read body error: %v", err)
		return err
	}

	v.wxtoken = string(getByBound(resBody, "window.wxtoken = \"", "\";"))
}

func (v *Voter) newappmsgvoteShow() error {
	url := v.getNewappmsgvoteShowUrl()
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("newappmsgvoteShow error: %v", err)
		return err
	}
	defer res.Body.Close()

	// TODO 看需要取什么内容 supervoteid是之前就有的
}

func (v *Voter) getNewappmsgvoteShowUrl() string {
	const newappmsgvoteShowUrl = "https://mp.weixin.qq.com/mp/newappmsgvote?action=show&__biz=%s&supervoteid=%s&uin=MTMwMzUxMjg3Mw%3D%3D&key=82438a29ddf26010ada8190c40aa8732323ae7c5941e6665f49b9e53af3ce594ebd25c63141a89301eebbf528329c02a1c7ac13e95a10b84f1c93ddf177ce4ee5b292cdd50d103fce5d0c369140dbee5&pass_ticket=WcK4v4itRVLRoKKVV0rGfjj4IWr2dK%252BXWGhasJO6LN6Ad1pRJMg1ShjC3mux%252BN8W&wxtoken=543112670&mid=2650886522&idx=1"
	return fmt.Sprintf(
		newappmsgvoteShowUrl,
		v.biz,
		v.supervoteid,
		v.uin,
	)
}

func (v *Voter) newappmsgvoteVote() error {
	// TODO
	// super_vote_item和super_vote_id是提前组装好的，因此只要encodeURIComponent即可
	url := "todo"
	res, err := v.Client.Post(url)
	if err != nil {
		log.Printf("newappmsgvoteVote error: %v", err)
		return err
	}
	defer res.Body.Close()
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
