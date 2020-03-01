package main

import (
	"fmt"
	"log"

	"github.com/JaesunJin/accounts"
)

func main() {
	account := accounts.NewAccount("jason")
	account.Deposit(10)
	err := account.Withdraw(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
