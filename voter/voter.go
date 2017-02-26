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
		// log.Printf("newappmsgvoteVote error: %v", err)
		return err
	}

	// main中根据vote结果确定是否要票数-1
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
	v.values.Set("supervoteid", v.Info.Supervoteid)
	v.values.Set("action", "show")
	log.Printf("newappmsgvoteShow values: %+v", v.values)

	url := getNewappmsgvoteShowUrl(v.values)
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("newappmsgvoteShow get error: %v", err)
		return err
	}
	defer res.Body.Close()

	// 这里不需要做什么，supervoteid是之前就有的
	// TODO 判断这个URL是否已经投过票。如果已投过，再投也会成功，因此需要在这里提前判断
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("newappmsgvoteShow read body error: %v", err)
		return err
	}
	log.Printf("newappmsgvoteShow resBody: %v", string(resBody))
	if bytes.Contains(resBody, []byte(`"selected":true`)) {
		log.Printf("newappmsgvoteShow already has vote")
		return errors.New("newappmsgvoteShow already has vote")
	}

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
	// item, err := json.Marshal(v.Info.Item)
	// if err != nil {
	// 	log.Printf("json.Marshal item error: %v", err)
	// 	return err
	// }
	// v.values.Set("json", string(item))
	v.values.Set("json", v.Info.Item)
	log.Printf("vote values: %v", v.values)
	log.Printf("newappmsgvoteVote formdata: %v", v.values.Encode())

	res, err := v.client.PostForm("https://mp.weixin.qq.com/mp/newappmsgvote", v.values)
	if err != nil {
		log.Printf("newappmsgvoteVote postform error: %v", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("newappmsgvoteVote read resBody error: %v", err)
		return err
	}
	// log.Printf("resBody: %v", string(resBody))

	resData := map[string]interface{}{}

	err = jsonUnmarshal(resBody, &resData)
	if err != nil || resData["base_resp"].(map[string]interface{})["ret"].(json.Number).String() != "0" {
		// 也是一种错误场景
		log.Printf("newappmsgvoteVote: resBody is invalid: %v", string(resBody))
		return errors.New("newappmsgvoteVote: resBody is invalid: " + string(resBody))
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
