package accounts

import (
	"errors"
	"fmt"
)

// Account comment
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your accunt
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of th account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner your Account
func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas:", a.Balance())
}
