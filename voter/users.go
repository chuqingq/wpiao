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
	IsAdmin  bool   `bson:"isadmin"`
}

// 第一步：如果带了正确的cookie，则成功，返回true，不返回结果
// 第二步：如果没带正确的cookie，且没带正确的用户名密码，则失败，返回false，返回结果
// 第三步：如果没带正确的cookie，带了正确的用户名密码，则成功，返回true，且设置了cookie，不返回结果
func UserLogin(w http.ResponseWriter, r *http.Request) *User {
	// log.Printf("UserLogin:")

	// 从数据库里读取所有用户名密码
	err := MgoFind("weipiao", "user", bson.M{}, &gUsers)
	if err != nil {
		log.Printf("MgoFind user error: %v", err)
		return nil
	}
	// log.Printf("load users: %v", len(gUsers))

	u := checkCookie(w, r)
	if u != nil {
		return u
	}

	// username、timestamp（单位是秒）、password、
	username := r.FormValue("username")
	timestamp := r.FormValue("timestamp")
	password := r.FormValue("password")
	// log.Printf("username: %v, password: %v, timestamp: %v", username, password, timestamp)

	u = check(username, password, timestamp)
	if u == nil {
		// log.Printf("check user error")
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

// 如果返回nil说明check失败，后续会判断用户名密码，因此这里不写入响应
func checkCookie(w http.ResponseWriter, r *http.Request) *User {
	// log.Printf("checkCookie: ")

	// 检查cookie是否合法、未过期
	usernameCookie, usernameCookieErr := r.Cookie("wp_username")
	passwordCookie, passwordCookieErr := r.Cookie("wp_password")
	timestampCookie, timestampCookieErr := r.Cookie("wp_timestamp")
	if usernameCookieErr != nil || passwordCookieErr != nil || timestampCookieErr != nil {
		log.Printf("checkCookie error: get cookie error: %v, %v, %v", usernameCookieErr, passwordCookieErr, timestampCookieErr)
		// w.Write([]byte(`{"ret":403,"msg":"cookie is invalid"}`))
		return nil
	}

	u := check(usernameCookie.Value, passwordCookie.Value, timestampCookie.Value)
	if u == nil {
		// log.Printf("check error")
		// w.Write([]byte(`{"ret":403,"msg":"check cookie error"}`))
		return nil
	}

	return u
}

func check(username, password, timestamp string) *User {
	// log.Printf("check: %v, %v, %v", username, password, timestamp)

	// timestamp是超时时间,单位是秒，password是salt+password做sha1后的结果
	if username == "" || password == "" || timestamp == "" {
		log.Printf("check error: param is empty: %v, %v, %v", username, password, timestamp)
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
	if byHex != password {
		log.Printf("password not match: %v, %v, %v, %v", username, password, timestamp, byHex)
		return nil
	}

	return u
}

func (user *User) Insert() error {
	return MgoInsert("weipiao", "user", user)
}

func (user *User) QueryAllUsers() ([]*User, error) {
	var users []*User
	err := MgoFind("weipiao", "user", bson.M{}, &users)
	if err != nil {
		log.Printf("MgoFind users error: %v", err)
		return nil, err
	}

	return users, nil
}

func (user *User) ChangePassword(pass string) error {
	return MgoUpdate("weipiao", "user", bson.M{"username": user.UserName}, bson.M{"$set": bson.M{"password": pass}})
}
