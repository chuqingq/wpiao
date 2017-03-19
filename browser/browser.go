package main

// go build -ldflags="-H windowsgui"

import (
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
)

const FILE_NAME = "D:\\abc.txt"

func main() {
	flag.Parse()
	voteUrl := flag.Arg(0)
	if voteUrl == "" {
		return
	}

	var content []byte
	var err error
	content, err = ioutil.ReadFile(FILE_NAME)
	if err != nil {
		content = make([]byte, 0)
	}

	str := string(content) + voteUrl + "\r\n"

	ioutil.WriteFile(FILE_NAME, []byte(str), 0666)

	// 请求通过http发送给
	reqUrl := `http://192.168.31.72:8080/api/vote?url=` + url.QueryEscape(voteUrl)
	http.Get(reqUrl)
}
