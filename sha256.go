package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var number int
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i:= 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			number++
		}
	}
	a := []string{}
	fmt.Println(a)
	fmt.Println("number is ", number)
}