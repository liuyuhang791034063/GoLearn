// @Date    : 2020-10-13
// @Author  : whiteyhliu
package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go Handler(conn)
	}
}

func Handler(c net.Conn)  {
	defer c.Close()
	log.Printf("%s has connected", c.RemoteAddr())
	for {
		args := make([]byte, 1024)
		length, err := c.Read(args)
		if err != nil {
			log.Fatal(err)
			return
		}
		op := strings.ToLower(string(args[:length-1]))

		switch op {
		case "get":
			get(c)
		case "send":
			send(c)
		case "ls":
			ls(c)
		case "cd":
			cd(c)
		default:
			_, err := c.Write([]byte("Please input operate"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func get(conn net.Conn)  {

}

func send(conn net.Conn)  {

}

func ls(conn net.Conn)  {

}

func cd(conn net.Conn)  {

}

