package main

import (
	// "encoding/hex"
	"log"
	"time"

	// log "github.com/cihub/seelog"
	"gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session /* mongo client */

func GetSession() *mgo.Session {
	return mongoSession
}

func InitMongo(mongoServer string) error {
	// pwbyte, err := hex.DecodeString(password)
	// if err != nil {
	// 	log.Printf("password string decode err: %v", err)
	// 	return err
	// }
	// pwDecry, err := DecryptImax(pwbyte)
	// if err != nil {
	// 	return err
	// }
	// pw := string(PKCS5Unpadding(pwDecry))
	// mongoSession, err = mgo.DialWithInfo((&mgo.DialInfo{Addrs: mongoServers, Username: userName, Password: pw, Database: "admin", Timeout: 10 * time.Second}))
	var err error
	mongoSession, err = mgo.Dial("mongodb://" + mongoServer)
	if err != nil {
		log.Printf("mongodb can't connect: %v", err)
		return err
	}
	// Optional. Switch the session to a monotonic behavior. secondary -> primary
	mongoSession.SetMode(mgo.Monotonic, true)
	mongoSession.SetSyncTimeout(time.Second * 5)
	return nil
}

// func InitMongo2(mongoServers []string, userName, password string) error {
// 	pwbyte, err := hex.DecodeString(password)
// 	if err != nil {
// 		log.Printf("password string decode err: %v", err)
// 		return err
// 	}
// 	pwDecry, err := DecryptImax(pwbyte)
// 	if err != nil {
// 		return err
// 	}
// 	pw := string(PKCS5Unpadding(pwDecry))
// 	mongoSession, err = mgo.DialWithInfo((&mgo.DialInfo{Addrs: mongoServers, Username: userName, Password: pw, Database: "admin", Timeout: 10 * time.Second}))
// 	//      mongoSession, err = mgo.Dial("mongodb://" + *mongoServer)
// 	if err != nil {
// 		log.Printf("mongodb can't connect: %v", err)
// 		return err
// 	}
// 	// Optional. Switch the session to a monotonic behavior. secondary -> primary
// 	mongoSession.SetMode(mgo.Monotonic, true)
// 	mongoSession.SetSyncTimeout(time.Second * 5)
// 	return nil
// }

func MgoFind(db string, collection string, query interface{}, result interface{}) error {
	session := mongoSession.Clone()
	defer session.Close()
	c := session.DB(db).C(collection)
	return c.Find(query).All(result)
}

func MgoInsert(db string, collection string, docs ...interface{}) error {
	session := mongoSession.Clone()
	defer session.Close()
	c := session.DB(db).C(collection)
	return c.Insert(docs...)
}

func MgoUpdate(db string, collection string, selector interface{}, update interface{}) error {
	session := mongoSession.Clone()
	defer session.Close()
	c := session.DB(db).C(collection)
	return c.Update(selector, update)
}

func MgoBulkUpdate(db string, collection string, pairs ...interface{}) error {
	session := mongoSession.Clone()
	defer session.Close()
	c := session.DB(db).C(collection)
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	res, err := bulk.Run()
	log.Printf("update result : %+v", res)
	return err
}
