package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	//balance variable to store state
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// Methods - convention to have receive variable as first letter of type - in this case "w". The type is defined by the "Wallet" struct in the test file.

// func (recivername ReceiverType) MethodName(arguments and type of arguement)
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// adding a pointer * to the Wallet struct
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//The var keyword allows us to define values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Question - why setting up point before the function name again?
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
