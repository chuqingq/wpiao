package main

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var gUsers Users

// type Users map[string]*User
type Users []*User

type User struct {
	UserName string `bson:"username"`
	Password string `bson:"password"`
}

// func init() {
// 	gUsers["user1"] = &User{
// 		UserName: "user1",
// 		Password: "user1",
// 	}
// }

// 第一步：如果带了正确的cookie，则成功，返回true，不返回结果
// 第二步：如果没带正确的cookie，且没带正确的用户名密码，则失败，返回false，返回结果
// 第三步：如果没带正确的cookie，带了正确的用户名密码，则成功，返回true，且设置了cookie，不返回结果
func UserLogin(w http.ResponseWriter, r *http.Request) *User {
	log.Printf("users Login:")

	// 从数据库里读取所有用户名密码
	err := MgoFind("weipiao", "user", bson.M{}, &gUsers)
	if err != nil {
		log.Printf("MgoFind user error: %v", err)
		return nil
	}

	u := checkCookie(w, r)
	if u != nil {
		return u
	}

	// TODO username、timestamp（单位是秒）、password、
	username := r.FormValue("username")
	log.Printf("username: %v", username)
	timestamp := r.FormValue("timestamp")
	log.Printf("timestamp: %v", timestamp)
	password := r.FormValue("password")
	log.Printf("password: %v", password)

	u = check(username, password, timestamp)
	if u == nil {
		log.Printf("check user error")
		w.Write([]byte(`{"ret":403,"msg":"username or password is invalid"}`))
		return nil
	}

	usernameCookie := &http.Cookie{
		Name:   "wp_username",
		Value:  username,
		MaxAge: 60 * 60 * 24, // 单位：秒。
	}
	http.SetCookie(w, usernameCookie)

	passwordCookie := &http.Cookie{
		Name:   "wp_password",
		Value:  password,
		MaxAge: 60 * 60 * 24, // 单位：秒。
	}
	http.SetCookie(w, passwordCookie)

	timestampCookie := &http.Cookie{
		Name:   "wp_timestamp",
		Value:  timestamp,
		MaxAge: 60 * 60 * 24, // 单位：秒。
	}
	http.SetCookie(w, timestampCookie)

	return u
}

func checkCookie(w http.ResponseWriter, r *http.Request) *User {
	log.Printf("checkCookie: ")

	// 检查cookie是否合法、未过期
	usernameCookie, usernameCookieErr := r.Cookie("wp_username")
	passwordCookie, passwordCookieErr := r.Cookie("wp_password")
	timestampCookie, timestampCookieErr := r.Cookie("wp_timestamp")
	if usernameCookieErr != nil || passwordCookieErr != nil || timestampCookieErr != nil {
		log.Printf("get cookie error")
		// w.Write([]byte(`{"ret":403,"msg":"cookie is invalid"}`))
		return nil
	}

	u := check(usernameCookie.Value, passwordCookie.Value, timestampCookie.Value)
	if u == nil {
		log.Printf("check error")
		// w.Write([]byte(`{"ret":403,"msg":"check cookie error"}`))
		return nil
	}

	return u
}

func check(username, password, timestamp string) *User {
	log.Printf("check: %v, %v, %v", username, password, timestamp)

	// timestamp是超时时间,单位是秒，password是salt+password做sha1后的结果
	if username == "" || password == "" || timestamp == "" {
		log.Printf("param is invalid")
		return nil
	}

	ts, err := strconv.Atoi(timestamp)
	if ts == 0 || err != nil {
		log.Printf("timestamp %v is invalid: %v", timestamp, err)
		return nil
	}

	t := time.Unix(int64(ts+5*60), 0)
	if t.Before(time.Now()) {
		log.Printf("timestamp %v expired", timestamp)
		return nil
	}

	// 获取密码
	// u := gUsers[username]
	var u *User
	for _, user := range gUsers {
		if user.UserName == username {
			u = user
			break
		}
	}

	if u == nil {
		log.Printf("user %v not found", username)
		return nil
	}

	// 计算password
	by := sha1.Sum([]byte(timestamp + u.Password))
	byHex := hex.EncodeToString(by[:])
	log.Printf("byHex: %v", byHex)
	if byHex != password {
		log.Printf("password %v not match")
		return nil
	}

	return u
}
