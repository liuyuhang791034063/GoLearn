package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now())
}
