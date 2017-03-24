package main

// 保存用户充值的支付订单号和金额

import (
	"errors"
	"log"

	// "github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

func SavePayOrder(order string, money float64) error {
	orders := []interface{}{}
	// 判断订单号是否重复
	err := MgoFind("weipiao", "payorder", bson.M{"order": order}, &orders)
	if err != nil {
		log.Printf("查询payorder失败：%v", err)
		return err
	}

	if len(orders) != 0 {
		log.Printf("订单号重复")
		return errors.New("订单号重复")
	}

	err = MgoInsert("weipiao", "payorder", bson.M{"order": order, "money": money, "handled": false})
	if err != nil {
		return err
	}
	return nil
}
