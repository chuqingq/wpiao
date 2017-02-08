package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	abool := robotgo.ShowAlert("test", "robotgo")
	if abool == 0 {
		fmt.Println("ok@@@", "ok")
	}

	title := robotgo.GetTitle()
	fmt.Println("title@@@", title)
}
