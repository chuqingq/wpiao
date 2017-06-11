package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Voter struct {
	url    string
	client *http.Client
	values url.Values
	Info   *Task
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
		// log.Printf("newappmsgvoteShow error: %v", err)
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
	// log.Printf("wxtoken: %v", v.values.Get("wxtoken"))
	if v.values.Get("wxtoken") == "" {
		log.Printf("get wxtoken error. maybe voteUrl expired")
		return errors.New("get wxtoken error. maybe voteUrl expired")
	}

	return nil
}

func (v *Voter) newappmsgvoteShow() error {
	// vi.supervoteid没有保存，因此需要从info中获取出supervoteid
	supervoteid := v.Info.Info["super_vote_id"].(int64)

	// 补齐参数
	v.values.Set("supervoteid", fmt.Sprintf("%v", supervoteid))
	v.values.Set("action", "show")
	// log.Printf("newappmsgvoteShow values: %+v", v.values)

	url := getNewappmsgvoteShowUrl(v.values)
	res, err := v.client.Get(url)
	if err != nil {
		log.Printf("newappmsgvoteShow get error: %v", err)
		return err
	}
	defer res.Body.Close()

	// 判断这个URL是否已经投过票。如果已投过，再投也会成功，因此需要在这里提前判断
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// log.Printf("newappmsgvoteShow read body error: %v", err)
		return err
	}
	// log.Printf("newappmsgvoteShow resBody: %v", string(resBody))
	if bytes.Contains(resBody, []byte(`"selected":true`)) {
		// log.Printf("newappmsgvoteShow error: 您已投票")
		return errors.New("newappmsgvoteShow error: 您已投票")
	}

	return nil
}

func getNewappmsgvoteShowUrl(values url.Values) string { // TODO 合并
	return "http://mp.weixin.qq.com/mp/newappmsgvote?" + values.Encode()
}

func (v *Voter) newappmsgvoteVote() error {
	v.values.Set("action", "vote")
	v.values.Set("f", "json")
	v.values.Set("json", v.Info.Item)
	// log.Printf("vote values: %v", v.values)
	// log.Printf("newappmsgvoteVote formdata: %v", v.values.Encode())

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
	log.Printf("resBody: %v", string(resBody))

	resData := map[string]interface{}{}

	err = jsonUnmarshal(resBody, &resData)
	if err != nil {
		log.Printf("newappmsgvoteVote jsonUnmarshal resBody error: %v", err)
		return errors.New("newappmsgvoteVote jsonUnmarshal resBody error")
	}

	retStr := resData["base_resp"].(map[string]interface{})["ret"].(json.Number).String()
	if retStr == "0" {
		log.Printf("vote success 投票成功\n")
		return nil
	}

	if retStr == "-6" {
		// log.Printf("newappmsgvoteVote vote error: 投票过于频繁，请稍后重试！")
		return errors.New("vote error: 投票过于频繁，请稍后重试！")
	}

	// 后面的错误可能都是任务本身的问题，例如需要关注、已过期，都把状态设置为fail，不再尝试
	v.Info.SetStatus("fail")

	if retStr == "-7" {
		// log.Printf("newappmsgvoteVote vote error: 关注公众号后才可以投票")
		return errors.New("vote error: 关注公众号后才可以投票")
	}

	// 其他失败
	// log.Printf("newappmsgvoteVote vote error: 投票失败: " + retStr)
	return errors.New("vote error: 投票失败: " + retStr)
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
