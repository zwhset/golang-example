package main

import (
	"fmt"
	"os"
	"github.com/labstack/gommon/log"
	"strings"
)

func main() {
	// Set env and Get a env
	err := os.Setenv("FLASK_ENV", "PRO")
	if err != nil {
		log.Fatal("Set Env Fail, %s\n", err)
	}
	fmt.Printf("FLASK_ENV: %s\n", os.Getenv("FLASK_ENV"))

	// Get all Env

	for _, env := range os.Environ() {
		//fmt.Printf(env)
		pair := strings.Split(env, "=") // 以=号分割环境变量
		fmt.Println(pair[0], "=", pair[1]) // 输出
	}

	// OUTPUT
	//FLASK_ENV: PRO
	//PATH = /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/usr/local/go/bin
	//SHELL = /bin/zsh
	//PAGER = less
	//LSCOLORS = Gxfxcxdxbxegedabagacad
	//OLDPWD = /Applications/GoLand.app/Contents/bin
	//GOPATH = /Users/zwhset/go
	//	VERSIONER_PYTHON_PREFER_32_BIT = no
	//USER = zwhset
	//GOROOT = /usr/local/go
	//	ZSH = /Users/zwhset/.oh-my-zsh
	//TMPDIR = /var/folders/k3/l1rymj0x6gqg9fk5b6h9f5l00000gn/T/
	//	SSH_AUTH_SOCK = /private/tmp/com.apple.launchd.MvOClaU29i/Listeners
	//XPC_FLAGS = 0x0
	//VERSIONER_PYTHON_VERSION = 2.7
	//__CF_USER_TEXT_ENCODING = 0x1F5:0x19:0x34
	//Apple_PubSub_Socket_Render = /private/tmp/com.apple.launchd.5iaGkyMany/Render
	//LESS = -R
	//LOGNAME = zwhset
	//LC_CTYPE =
	//	XPC_SERVICE_NAME = com.jetbrains.goland.7548
	//PWD = /Users/zwhset/code/github/golang-example
	//HOME = /Users/zwhset
	//FLASK_ENV = PRO
}
