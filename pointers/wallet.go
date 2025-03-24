package main

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amout Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p\n", &w.balance)
	w.balance += amout
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amout Bitcoin) error {
	if amout > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amout
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
