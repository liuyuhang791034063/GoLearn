package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	maps := make(map[string]int)
	var lines int

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		str := input.Text()
		maps[str]++
		lines++
	}

	fmt.Printf("The text have %d lines\n", lines)
	for k, v := range maps {
		fmt.Printf("The %s find number is %d\n", k, v)
	}
}
