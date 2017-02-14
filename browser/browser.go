package main

import (
	"flag"
	"io/ioutil"
)

const FILE_NAME = "D:\\abc.txt"

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		return
	}

	var content []byte
	var err error
	content, err = ioutil.ReadFile(FILE_NAME)
	if err != nil {
		content = make([]byte,0)
	}

	str := string(content)
	str = str + url + "\n"

	ioutil.WriteFile(FILE_NAME, []byte(str), 0666)
}
