package main

import (
	"ch12/sexpr"
	"fmt"
)

func main()  {
	test := struct {
		Name 	string
		time 	struct{
			Year 	int
			Day 	int
			Test 	int
		}
		test 	int
	}{"123", struct {
		Year int
		Day  int
		Test int
	}{Year: 123, Day: 123, Test: 123}, 123}
	res, _ := sexpr.Marshal(test)
	var newTest interface{}
	err := sexpr.Unmarshal(res, newTest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newTest)
}
