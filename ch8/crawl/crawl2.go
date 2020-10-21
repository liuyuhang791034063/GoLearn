// @Date    : 2020-10-21
// @Author  : whiteyhliu
package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Link struct{
	url string
	deep int
}

func crawl(url string, deep int) []Link {
	if deep > 4 {
		return nil
	}
	fmt.Println(url)
	return []Link{
		{string(rand.Intn(100)), deep},
		{string(rand.Intn(100)), deep},
		{string(rand.Intn(100)), deep},
	}
}

func main() {
	worklist := make(chan []Link)
	unseenLinks := make(chan Link)
	var n int

	n++

	go func() { worklist <- []Link{{os.Args[1], 1}}}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.url, link.deep+1)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for lists := range worklist {
		for _, link := range lists {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

