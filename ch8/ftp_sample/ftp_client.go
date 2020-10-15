// @Date    : 2020-10-13
// @Author  : whiteyhliu
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Has connected %s\n", conn.LocalAddr())
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\r')
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = conn.Write([]byte(str))
		if err!= nil {
			log.Fatal(err)
		} else {
			fmt.Printf("send success, message is %s\n", str)
		}

		rMsg := make([]byte, 1024)
		_, err = conn.Read(rMsg)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("receive message is %s\n",string(rMsg))
	}
}