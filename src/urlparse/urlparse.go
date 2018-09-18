package main

import (
	"net/url"
	"github.com/labstack/gommon/log"
	"fmt"
	"net"
)

func main() {
	//dbInfo := "postgres://user:pass@host.com:5432/path?k=v#f"
	dbInfo := "https://baike.baidu.com/item/SHA1/8812671?fr=aladdin&s=good"

	data, err := url.Parse(dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	scheme := data.Scheme
	user := data.User
	host := data.Host
	password, _ := data.User.Password()
	fmt.Printf("Scheme: %s\nHOST: %s\nUSER: %s\nPASSWORD: %s\n\n", scheme, host, user, password)

	host, port, _ := net.SplitHostPort(host) // Split host and port
	path := data.Path
	fmt.Printf("HOST: %s\nPORT: %s\nPATH: %s\n\n", host, port, path)
	fmt.Println(data.RawQuery)

	ks, _ := url.ParseQuery(data.RawQuery)
	fmt.Println(ks)

	// OUTPUT:
	//Scheme: https
	//HOST: baike.baidu.com
	//USER:
	//PASSWORD:
	//
	//HOST:
	//PORT:
	//PATH: /item/SHA1/8812671
	//
	//	fr=aladdin&s=good
	//	map[fr:[aladdin] s:[good]]
}
