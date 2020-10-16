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
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Has connected %s\n", conn.LocalAddr())

	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
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
		fmt.Printf("%s", rMsg)
	}
}