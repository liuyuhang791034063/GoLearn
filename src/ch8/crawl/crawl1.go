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
	var n int

	n++

	go func() { worklist <- []Link{{os.Args[1], 1}}}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		links := <- worklist
		for _, link := range links {
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link Link) {
					worklist <- crawl(link.url, link.deep+1)
				}(link)
			}
		}
	}
}

