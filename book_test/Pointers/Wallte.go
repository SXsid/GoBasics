package pointers

import "fmt"

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

func (wallet *Wallet) Withdraw(withdrawAmount Bitcoin) {
	wallet.amount -= withdrawAmount
}
