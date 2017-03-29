package main

import (
	"errors"
	"log"
	"time"

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

// main调用此方法首次下发任务
func RunnersDispatchTask(task *Task) error {
	// 先扣除费用。可能失败
	user := gUsers.GetUserByName(task.User)
	if user.Balance < float64(task.Votes)*task.Price {
		log.Println("用户 %s 余额不足: %f < %d*%f", task.User, user.Balance, task.Votes, task.Price)
		return errors.New("任务下发失败：账户余额不足")
	}
	err := user.SetBalance(user.Balance - float64(task.Votes)*task.Price)
	if err != nil {
		log.Println("计费失败: %v", err)
		return errors.New("任务下发失败：内部错误：计费失败")
	}

	doDispatchTask(task)
	return nil
}

// RunnersDispatchTask和NotifyTaskFinish调用此方法分配给runner
func doDispatchTask(task *Task) {
	// 设置执行task的runner个数，标记何时结束
	task.SetRunnerCount(len(gRunners))

	// TODO 把任务按比例分给runner。目前是平均，还不支持按比例
	//  根据当前curRunner和curIndex分配runner
	votes := int(task.Votes - task.CurVotes)
	vote1 := votes / len(gRunners)
	vote2 := votes % len(gRunners)
	index := 0
	for _, r := range gRunners {
		count := vote1
		if index < vote2 {
			count += 1
		}
		// 经过验证，这种算法基本平均
		r.DispatchTask(task, count)
		log.Printf("doDispatchTask: runner: %v, task: %v, count: %v", r.Name, task.Id, count)
		index++
	}
}

func (r *Runner) NotifyTaskFinish(task *Task) {
	// 在数据库中标记该任务又结束了一个runner
	// runnerCount := task.DecrRunnerCount() // 直接返回当前正在执行的runner数
	err := task.DecrRunnerCount()
	if err != nil {
		log.Printf("NotifyTaskFinish task.DecrRunnerCount() error: %v", err)
		return
	}
	// 如果还有别的runner未结束，则继续等待，不做动作
	if task.RunnerCount > 0 {
		log.Printf("该任务还有runner在运行，等待。。。")
		return
	}
	log.Printf("该任务runner均结束")

	// TODO 如果所有runner都结束了，判断是否要重新下发任务来补充差额
	// finish := task.Votes <= task.CurVotes // 这种方式是不停的跑，知道票数OK
	finish := true // 这种方式是只下发一次
	if !finish {
		// 需要重新下发
		doDispatchTask(task)
		return
	}

	time.Sleep(5 * time.Second)

	// 如果不补充差额，则任务结束，返回差额
	task.SetFinishTime(time.Now())
	task.SetStatus("finished")
	if task.Votes <= task.CurVotes {
		log.Printf("任务已完成，无需退款: %v < %v", task.Votes, task.CurVotes)
		return
	}
	// 如果补充差额，则通过DoDispatchTask(task)来补充差额
	user := gUsers.GetUserByName(task.User)
	tuikuan := float64(task.Votes-task.CurVotes) * task.Price
	log.Printf("需退款：%v", tuikuan)
	err = user.SetBalance(user.Balance + tuikuan)
	if err != nil {
		log.Println("退款失败: %v", err)
		return
	}
}

func (r *Runner) DispatchTask(task *Task, votes int) {
	// votes := task.Votes - task.CurVotes
	// if votes <= 0 {
	// 	return
	// }

	req := map[string]interface{}{}
	req["cmd"] = "vote"
	req["url"] = task.Url
	// TODO pc那边需要保存这个taskid，且在任务结束时返回
	req["taskid"] = task.Id
	req["votes"] = votes

	err := r.Conn.WriteJSON(req)
	if err != nil {
		log.Printf("ws.WriteJSON error: %v", err)
	}
}
