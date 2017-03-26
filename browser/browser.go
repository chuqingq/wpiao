package main

// go build -ldflags="-H windowsgui"

import (
	"flag"
	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// const FILE_NAME = "D:\\abc.txt"

func init() {
	file, err := os.OpenFile("browser.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("open voter.log error: %v", err)
	}
	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}

func main() {
	flag.Parse()
	voteUrl := flag.Arg(0)
	if voteUrl == "" {
		return
	}

	// var content []byte
	// var err error
	// content, err = ioutil.ReadFile(FILE_NAME)
	// if err != nil {
	// 	content = make([]byte, 0)
	// }

	// str := string(content) + voteUrl + "\r\n"

	// ioutil.WriteFile(FILE_NAME, []byte(str), 0666)
	log.Printf("voteUrl: %v", voteUrl)

	// 请求通过http发送给
	reqUrl := `http://192.168.31.72:8080/api/vote?url=` + url.QueryEscape(voteUrl)
	resp, err := http.Get(reqUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("http.Get error: %v, %v", err, resp.StatusCode)
		return
	}
	log.Printf("http.Get success")
}
