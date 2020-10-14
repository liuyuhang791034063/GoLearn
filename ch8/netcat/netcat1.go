// @Date    : 2020-10-13
// @Author  : whiteyhliu
package main

import (
	"io"
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
	var f io.Writer
	f, err = os.OpenFile("test.log", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(f, conn)
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
