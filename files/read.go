package main

import (
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"fmt"
)

func main() {

	filename := "/Users/zwhset/code/qianbao/Dockerfile"
	// use ioutil.ReadFile
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Open File Fail, %s", err)
	}
	fmt.Print(string(bytes)) // bytes to string
}