// @Date    : 2020-10-12
// @Author  : whiteyhliu
package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handlerConn(conn)
	}
}

