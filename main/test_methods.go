package main

import (
	"ch12/methods"
	"strings"
	"time"
)

func main() {
	methods.Print(time.Hour)

	methods.Print(new(strings.Replacer))
}