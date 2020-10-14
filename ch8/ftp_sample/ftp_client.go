// @Date    : 2020-10-13
// @Author  : whiteyhliu
package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		str := os.Stdin
	}
}