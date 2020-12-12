// @Date    : 2020-10-12
// @Author  : whiteyhliu
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonaci(%d) = %d\n", n, fibN)
}

func spinner(delpy time.Duration)  {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delpy)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return 2
	}
	return fib(x-1) + fib(x-2)
}