// @Date    : 2020-10-28
// @Author  : whiteyhliu
package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

type reWithdraw struct {
	re chan bool
	amount int
}
var re = make(chan bool)
var withdraws = make(chan reWithdraw)

func Deposit(amount int)  {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int)  {
	withdraws <- struct {
		re     chan bool
		amount int
	}{re: re, amount: amount}
	for flag := range re {
		if flag{
			fmt.Println("done")
			return
		}else {
			fmt.Println("dont")
			return
		}
	}
}

func teller()  {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			continue
		case withdraw := <-withdraws:
			if balance >= withdraw.amount {
				balance -= withdraw.amount
				withdraw.re <- true
			}else {
				withdraw.re <- false
			}
		}
	}
}

func init()  {
	go teller()
}

