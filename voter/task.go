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

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// // key -> task
// type Tasks map[string]*Task

// func (vis Tasks) Get(key string) *Task {
// 	return vis[key]
// }

// func (vis Tasks) Set(key string, vi *Task) {
// 	vis[key] = vi
// }

// func (vis Tasks) Del(key string) {
// 	delete(vis, key)
// }

type Task struct {
	Id          bson.ObjectId          `bson:"_id"`
	Key         string                 `bson:"key"`    // 可以唯一标识一个投票的 TODO
	Status      string                 `bson:"status"` // 任务状态：prepare,doing,finished
	Url         string                 `bson:"url"`    // 短URL
	Supervoteid string                 `bson:"supervoteid"`
	Info        map[string]interface{} `bson:"info"` // 投票信息。包括活动标题、到期时间、投票对象等
	// Info     string `bson:"info"`
	Item        string  `bson:"item"`  // Item        map[string]interface{} `bson:"item"`  // 投的对象
	User        string  `bson:"user"`  // 下发任务的用户名
	Votes       uint64  `bson:"votes"` // 票数
	Price       float64 `bson:"price"` // 单价，单位是元/票
	Speed       uint64  `bson:"speed"` // TODO 暂未使用。每分钟的票数
	CurVotes    uint64  `bson:"curvotes"`
	RunnerCount int     `bson:""runnercount` // 在运行的runner数量
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

func NewTask(shortOrLongUrl string) (*Task, error) {
	log.Printf("NewTask inputUrl: %v", shortOrLongUrl)

	// 设置cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("cookiejar.New() error: %v", err)
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	vi := &Task{
		Url:    strings.Replace(shortOrLongUrl, "https:", "http:", 1),
		Status: "prepare",
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

	// voteInfoStr := string(getByBound(resBody, []byte(`var voteInfo=`), []byte(`;`)))
	voteInfoBytes := getByBound(resBody, []byte(`var voteInfo=`), []byte(`;`))
	log.Printf("voteInfoStr: %v ...", string(voteInfoBytes[:60]))
	vi.Info = make(map[string]interface{})
	// err = json.Unmarshal([]byte(voteInfoStr), &vi.Info)
	err = jsonUnmarshal(voteInfoBytes, &vi.Info)
	// log.Printf("info: %v", vi.Info)
	if err != nil {
		log.Printf("json.Unmarshal voteInfo error: %v", err)
		return nil, err
	}

	// TODO 保存key，传到前端。后面下发任务时再传回来
	// vi.Info["key"] = vi.Key

	return vi, nil
}

func (vi *Task) NewVoter(voteUrl string) (*Voter, error) {
	log.Printf("NewVoter():")
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
		client: &http.Client{
			Jar: jar,
		},
		values: u.Query(),
		Info:   vi,
	}, nil
}

func (vi *Task) Insert() error {
	return MgoInsert("weipiao", "task", vi)
}

// 提交任务
func (vi *Task) Submit() error {
	// update votes/item等字段 TODO
	return MgoInsert("weipiao", "task", vi)
}

func QueryTasksByUser(username string) ([]*Task, error) {
	var task []*Task
	err := MgoFind("weipiao", "task", bson.M{"user": username}, &task)
	if err != nil {
		log.Printf("MgoFind(task) error: %v", err)
		return nil, err
	}

	return task, nil
}

func QueryTaskById(taskId string) (*Task, error) {
	log.Printf("QueryTaskById(): %v", taskId)
	var tasks []*Task
	err := MgoFind("weipiao", "task", bson.M{"_id": bson.ObjectIdHex(taskId)}, &tasks)
	if err != nil {
		log.Printf("MgoFind(task) error: %v", err)
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("task not found by key")
	}

	return tasks[0], nil
}

func QueryTaskByKey(key string) (*Task, error) {
	var task []*Task
	err := MgoFind("weipiao", "task", bson.M{"key": key, "status": "doing"}, &task)
	if err != nil {
		log.Printf("MgoFind(task) error: %v", err)
		return nil, err
	}

	if len(task) == 0 {
		return nil, errors.New("task not found by key")
	}

	vi := task[0]

	// 领任务时票数+1。后面如果投失败，则-1
	vi.IncrVotes()

	return vi, nil
}

func QueryTaskBySuperVoteId(supervoteid string) (*Task, error) {
	var task []*Task
	err := MgoFind("weipiao", "task", bson.M{"supervoteid": supervoteid}, &task)
	if err != nil {
		log.Printf("MgoFind(task) error: %v", err)
		return nil, err
	}

	if len(task) == 0 {
		return nil, errors.New("task not found: supervoteid: " + supervoteid)
	}

	return task[0], nil
}

func (vi *Task) IncrVotes() error {
	log.Printf("vi.IncrVotes")
	err := MgoUpdate("weipiao", "task", bson.M{"_id": vi.Id}, bson.M{"$inc": bson.M{"curvotes": 1}})
	if err != nil {
		log.Printf("mgoupdate incr curvotes error: %v", err)
		return err
	}

	vi.CurVotes += 1
	if vi.CurVotes < vi.Votes {
		return nil
	}

	vi.Status = "success"
	log.Printf("task status: success, %v", vi.Id)
	return MgoUpdate("weipiao", "task", bson.M{"_id": vi.Id}, bson.M{"$set": bson.M{"status": vi.Status}})
	// return MgoUpdate("weipiao", "task", bson.M{"key": vi.Key}, bson.M{"$set": bson.M{"curvotes": vi.CurVotes, "status": vi.Status}})
}

func (vi *Task) DecrVotes() error {
	log.Printf("vi.DecrVotes")
	err := MgoUpdate("weipiao", "task", bson.M{"_id": vi.Id}, bson.M{"$inc": bson.M{"curvotes": -1}})
	if err != nil {
		log.Printf("mgoupdate decr curvotes error: %v", err)
		return err
	}

	vi.CurVotes -= 1
	if vi.CurVotes >= vi.Votes {
		return nil
	}

	vi.Status = "doing"
	log.Printf("task status: doing, %v", vi.Id)
	return MgoUpdate("weipiao", "task", bson.M{"_id": vi.Id}, bson.M{"$set": bson.M{"status": vi.Status}})
	// return MgoUpdate("weipiao", "task", bson.M{"key": vi.Key}, bson.M{"$set": bson.M{"curvotes": vi.CurVotes, "status": vi.Status}})
}

func (vi *Task) SetStatus(status string) error {
	vi.Status = status
	return MgoUpdate("weipiao", "task", bson.M{"key": vi.Key}, bson.M{"$set": bson.M{"status": vi.Status}})
}

func (task *Task) DecrRunnerCount() error {
	// 原子-1并返回内容
	session := mongoSession.Clone()
	defer session.Close()
	c := session.DB("weipiao").C("task")

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"runnercount": -1}},
		ReturnNew: true,
	}
	// var task2 Task
	_, err := c.Find(bson.M{"_id": bson.ObjectId(task.Id)}).Apply(change, &task)
	if err != nil {
		log.Printf("DecrRunnerCount change error: %v", err)
		return err
	}
	log.Printf("new runnercount: %v", task.RunnerCount)
	// err := MgoUpdate("weipiao", "task", bson.M{"_id": bson.ObjectId(task.Id)}, bson.M{"$inc": bson.M{"runnercount": -1}})
	// if err != nil {
	// 	log.Printf("DecrRunnerCount error: %v", err)
	// 	return -1
	// }
	// task2, err := QueryTaskById(task.Id.Hex())
	// if err != nil {
	// 	log.Printf("QueryTaskById error: %v", err)
	// 	return -1
	// }
	return nil
}

func (task *Task) SetRunnerCount(count int) error {
	return MgoUpdate("weipiao", "task", bson.M{"_id": task.Id}, bson.M{"$set": bson.M{"runnercount": count}})
}

func jsonUnmarshal(data []byte, v interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

// // 一个PC完成后根据url做调整
// func TaskDispatch(voteUrl string) error {
// 	log.Printf("TaskDispatch: url: %v", voteUrl)

// 	var tasks []*Task
// 	err := MgoFind("weipiao", "task", bson.M{"url": voteUrl, "status": "doing"}, &tasks)
// 	if err != nil {
// 		log.Printf("MgoFind(task) error: %v", err)
// 		return err
// 	}

// 	if len(tasks) == 0 {
// 		return errors.New("task not found: url: " + voteUrl)
// 	}

// 	task := tasks[0]
// 	if task.Status != "doing" {
// 		return nil
// 	}

// 	if task.Status == "doing" {
// 		r := GetFreeRunner(task.Key) // TODO 还是要保证db.task中记录是有key的
// 		if r == nil {
// 			log.Printf("TaskDispatch: no free runner: %v", task.Key)
// 			return errors.New("TaskDispatch: no free runner")
// 		}
// 		r.DispatchTask(task)
// 	}

// 	return nil
// }
