package pointers

import (
	"errors"
	"fmt"
)

type Wallet struct {
	amount Bitcoin
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(depositAmount Bitcoin) {
	wallet.amount += depositAmount

}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(withdrawAmount Bitcoin) error {

	if withdrawAmount > wallet.amount {
		return ErrInsufficientFunds
	}
	wallet.amount -= withdrawAmount
	return nil
}
