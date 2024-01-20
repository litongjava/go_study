package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fulldesc := "root@192.168.3.9"
	fmt.Println(fulldesc)

	fulldescArray := strings.Split(fulldesc, "@")
	username := fulldescArray[0]
	hostAndPort := strings.Split(fulldescArray[1], ":")
	host := hostAndPort[0]
	port := ""
	if len(hostAndPort) > 1 {
		port = hostAndPort[1]
	} else {
		port = "22"
	}

	log.Println("", host, port, username)
}
