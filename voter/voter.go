package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
	// TODO 如果投票失败，还需要再增加投票

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
	// v.Info.CurVotes += 1
	v.Info.IncrVotes()
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

	resData := map[string]interface{}{}

	err = jsonUnmarshal(resBody, &resData)
	if err != nil || resData["base_resp"].(map[string]interface{})["ret"].(json.Number).String() != "0" {
		// 也是一种错误场景
		log.Printf("vote error: %v", string(resBody))
		return errors.New("vote error: vote response is invalid: " + string(resBody))
	}

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
