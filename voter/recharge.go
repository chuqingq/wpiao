package main

// 保存用户充值的支付订单号和金额

import (
	"errors"
	"log"

	// "github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// 管理员录入充值订单号和金额
func RecordRechargeOrder(order string, money float64) error {
	orders := []interface{}{}
	// 判断订单号是否重复
	err := MgoFind("weipiao", "recharge", bson.M{"order": order}, &orders)
	if err != nil {
		log.Printf("查询recharge order失败：%v", err)
		return err
	}

	if len(orders) != 0 {
		log.Printf("订单号重复")
		return errors.New("订单号重复")
	}

	err = MgoInsert("weipiao", "recharge", bson.M{"order": order, "money": money, "handled": false})
	if err != nil {
		return err
	}
	return nil
}
