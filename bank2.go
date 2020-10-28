// @Date    : 2020-10-28
// @Author  : whiteyhliu
package main

import (
	"bank"
	"fmt"
)

func main() {
	fmt.Println(bank.Balance())
	bank.Deposit(123)
	fmt.Println(bank.Balance())
	bank.Withdraw(12)
	fmt.Println(bank.Balance())
}
