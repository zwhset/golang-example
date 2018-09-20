package main

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/gommon/log"
)

func main() {
	data := "zwhset"

	// string encode
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("String Encode: %s => %s\n", data, string(sEnc))

	// string decode
	decData := "endoc2V0"
	sDec, err := base64.StdEncoding.DecodeString(decData)
	if err != nil {
		log.Fatal("String: %s Decode Base64 Fail, %s", decData, err)
	}
	fmt.Printf("String Decode: %s => %s\n", decData, string(sDec))

	urlData := "http://www.weiops.com/search?key=golang&username=zwhset"
	// url encode
	uEnc := base64.URLEncoding.EncodeToString([]byte(urlData))
	fmt.Printf("URL Encode: %s => %s\n", urlData, string(uEnc))

	// url decode
	urlDecData := "aHR0cDovL3d3dy53ZWlvcHMuY29tL3NlYXJjaD9rZXk9Z29sYW5nJnVzZXJuYW1lPXp3aHNldA=="
	uDec, err := base64.URLEncoding.DecodeString(urlDecData)
	if err != nil {
		log.Fatal("String: %s URLDecode Base64 Fail, %s", uDec, err)
	}
	fmt.Printf("URL Decode: %s => %s\n", urlDecData, string(uDec))

	// OUTPUT:
	//String Encode: zwhset => endoc2V0
	//String Decode: endoc2V0 => zwhset
	//URL Encode: http://www.weiops.com/search?key=golang&username=zwhset => aHR0cDovL3d3dy53ZWlvcHMuY29tL3NlYXJjaD9rZXk9Z29sYW5nJnVzZXJuYW1lPXp3aHNldA==
	//URL Decode: aHR0cDovL3d3dy53ZWlvcHMuY29tL3NlYXJjaD9rZXk9Z29sYW5nJnVzZXJuYW1lPXp3aHNldA== => http://www.weiops.com/search?key=golang&username=zwhset
}
