package main

import (
	"os"

	"github.com/bubbly-hb/blogSystem-gin/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
