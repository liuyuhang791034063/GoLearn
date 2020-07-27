package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now())
}
