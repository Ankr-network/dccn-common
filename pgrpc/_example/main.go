package main

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	go server()
	client()
}

func genHook(prefix string) func(conn *net.Conn, err error) {
	return func(conn *net.Conn, err error) {
		if err != nil {
			log.Println(prefix, err)
		}
	}
}
