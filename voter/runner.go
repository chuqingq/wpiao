package main

import (
	"log"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Runner struct {
	Name         string `json:"name"`
	AccountCount int    `json:"accounts"`
	Conn         *websocket.Conn
	Addr         string `json:"addr"`
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
	log.Printf("taskrunnerGetFreeOne: key: %v", key)

	var taskRunner []*TaskRunner
	err := MgoFind("weipiao", "taskrunner", bson.M{"key": key}, &taskRunner)
	if err != nil {
		log.Printf("mgofind taskrunner error: %v", err)
		return nil
	}

	// 遍历全局gPC，第一个不在taskKey中的，返回 TODO 加锁
	for runner, _ := range gRunners {
		found := false
		for _, runnerDB := range taskRunner {
			if runner == runnerDB.Runner {
				found = true
				break
			}
		}

		if found {
			continue
		}

		// 更新数据库，占用空闲的PC
		err := MgoInsert("weipiao", "taskrunner", bson.M{"key": key, "runner": runner})
		if err != nil {
			log.Printf("mgoinsert taskrunner error: %v", err)
		}

		log.Printf("GetFreeRunner: key:%v, runner: %+v", key, gRunners[runner])
		return gRunners[runner]
	}

	log.Printf("taskrunnerGetFreeOne: no free runner for key: %v", key)
	return nil
}

func (r *Runner) DispatchTask(task *Task) {
	votes := task.Votes - task.CurVotes
	if votes <= 0 {
		return
	}

	// 检查用户余额是否足够
	user := gUsers.GetUserByName(task.User)
	if user.Balance < float64(task.Votes)*task.Price {
		log.Println("用户 %s 余额不足: %f < %d*%f", task.User, user.Balance, task.Votes, task.Price)
		return
	}
	err := user.SetBalance(user.Balance - float64(task.Votes)*task.Price)
	if err != nil {
		log.Println("计费失败: %v", err)
		return
	}

	req := map[string]interface{}{}
	req["cmd"] = "vote"
	req["url"] = task.Url
	req["votes"] = votes

	err = r.Conn.WriteJSON(req)
	if err != nil {
		log.Printf("ws.WriteJSON error: %v", err)
	}
}
