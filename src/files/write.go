package main

import (
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"os"
	"bufio"
	"fmt"
)

func main() {
	filename := "./gofile"
	filename2 := "./gofile2.txt"

	// write file
	data := []byte("hello\tgo\n")
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Write File Fail, %s", err)
	}

	// create file
	fd, err := os.Create(filename2)
	if err != nil {
		log.Fatal("Create File Fail, %s", err)
	}
	fd.WriteString("Good Boy for zwhset.\n") // write string
	fd.Sync() // Sync to Disk
	defer fd.Close()

	w := bufio.NewWriter(fd)
	n4, err := w.WriteString("This is ok?")
	if err != nil {
		log.Fatal("Write Content File Fail, %s", err)
	}
	fmt.Printf("File Object: %d", n4)
}
