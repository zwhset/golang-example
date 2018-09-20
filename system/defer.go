package main

import (
	"os"
	"github.com/labstack/gommon/log"
)

func main() {
	filename := "/tmp/defer.txt"

	fd, err := os.Create(filename)
	if err != nil {
		log.Fatal("Open File Fail, %s\n", err)
	}

	defer fd.Close() // 不管函数执行成功与否都会关闭文件描述符

	fd.WriteString("Test defer to close file.")
	fd.Sync()
}

func createFile(filename string) error {
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer fd.Close()
	fd.Sync()
}

func writeFile(filename string, content []byte) (n int, err error) {
	fd, err := os.Open(filename)

	defer fd.Close()
	if err != nil {
		return 0, err
	}
	n, err = fd.Write(content)
	fd.Sync()
	return n, err
}

func closeFile(fd *os.File)  {
	fd.Close()
	return
}
