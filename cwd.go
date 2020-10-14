package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Printf(cwd)
}
