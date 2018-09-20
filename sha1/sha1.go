package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"github.com/labstack/gommon/log"
)

func main() {
	// Check String Sha1
	s := "sha1 this string" // string
	h := sha1.New() // creaet a New object
	h.Write([]byte(s))
	bs := h.Sum(nil) // bytes sha1

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// Check File Sha1
	filename := "gofile"
	fileBs := fileSha1(filename)
	fmt.Printf("%s => %x\n", filename, fileBs)

	// Shell OUTPUT:
	//# shasum gofile
	//ca24369a08c2e01c073c20fc236d2fd0d28388fe  gofile

	// GOLANG OUTPUT:
	//sha1 this string
	//cf23df2207d99a74fbe169e3eba035e633b65d94
	//gofile => ca24369a08c2e01c073c20fc236d2fd0d28388fe

}

func fileSha1(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Open File %s Fail.", filename)
	}

	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	return bs
}