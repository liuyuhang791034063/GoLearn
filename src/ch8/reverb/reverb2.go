// @Date    : 2020-10-20
// @Author  : whiteyhliu
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(conn net.Conn, shout string, delay time.Duration) {
	wg.Done()
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	var message = make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			wg.Add(1)
			select {
			case <-time.After(10 * time.Second):
				conn.Close()
			case mes := <-message:
				go echo(conn, mes, 1*time.Second)
			}
		}
	}()
	for input.Scan() {
		text := input.Text()
		message <- text
	}
	wg.Wait()
	conn.Close()
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
