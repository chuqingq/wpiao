package main

import (
	"log"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Runner struct {
	Name         string
	AccountCount int
	Conn         *websocket.Conn
}

// 全局变量。有哪些PC已连上来
var gRunners = map[string]*Runner{}

// 数据库weipiao集合taskrunner中记录的task和runner对应关系
type TaskRunner struct {
	Key    string `bson:"key"`
	Runner string `bson:"runner"`
}

// 获取task未使用的pc
func GetFreeRunner(key string) *Runner {
	log.Printf("TaskPCGetFreeOne: key: %v", key)

	var taskRunner []*TaskRunner
	err := MgoFind("weipiao", "taskrunner", bson.M{"key": key}, &taskRunner)
	if err != nil {
		log.Printf("mgofind taskrunner error: %v", err)
		return nil
	}

	// 遍历全局gPC，第一个不在taskKey中的，返回 TODO 加锁
	for runner, _ := range gRunners {
		for _, runnerDB := range taskRunner {
			if runner == runnerDB.Key {
				continue
			}
		}

		// 更新数据库，占用空闲的PC
		err := MgoInsert("weipiao", "taskpc", bson.M{"key": key, "runner": runner})
		if err != nil {
			log.Printf("mgoinsert taskpc error: %v", err)
		}
		return gRunners[runner]
	}

	log.Printf("TaskPCGetFreeOne: no free runner for key: %v", key)
	return nil
}

func (r *Runner) DispatchTask(task *Task) {
	votes := task.Votes - task.CurVotes
	if votes <= 0 {
		return
	}

	req := map[string]interface{}{}
	req["cmd"] = "vote"
	req["url"] = task.Url
	req["votes"] = votes

	err := r.Conn.WriteJSON(req)
	if err != nil {
		log.Printf("ws.WriteJSON error: %v", err)
	}
}
