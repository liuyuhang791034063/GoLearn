// @Date    : 2020-10-13
// @Author  : whiteyhliu
package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup = sync.WaitGroup{}
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	defer listen.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		wg.Add(1)
		go Handler(conn)
	}
	wg.Wait()
}

func Handler(c net.Conn)  {
	defer wg.Done()
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
		op = strings.Split(op, "\r")[0]
		switch op {
		case "get":
			get(c)
		case "send":
			send(c)
		case "ls":
			message := ls()
			_, err := c.Write([]byte(message))
			if err != nil {
				log.Fatal(err)
			}
		case "cd":
			cd(c)
		case "close":
			c.Close()
			return
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

func ls() string {
	var res string
	pwd, err := os.Getwd()
	if err != nil {
		message := "os.Getwd function is err, message is " + err.Error()
		log.Fatal(message)
		return message
	}
	res += pwd + "\n"

	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for flag, file := range fileInfoList {
		if flag != 0 && flag % 5 == 0 {
			res += "\n"
		}
		res += file.Name() + "\t"
	}
	return res
}

func cd(conn net.Conn)  {

}

