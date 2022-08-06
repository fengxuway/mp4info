package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fengxuway/mp4info"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	duration, err := mp4info.GetMP4Duration(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(filepath.Base(os.Args[1]), duration)
}
