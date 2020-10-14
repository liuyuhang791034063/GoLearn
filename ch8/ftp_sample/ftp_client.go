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
	defer conn.Close()
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
			fmt.Printf("send success, message is %s", str)
		}
	}
}