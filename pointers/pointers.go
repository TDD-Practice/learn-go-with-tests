package pointers

import "errors"

type Wallet struct {
	balance float32
}

var ErrInsufficientFunds = errors.New("not enough founds to withraw the amount")

func (w *Wallet) Deposit(amount float32) {
	w.balance += amount
}

func (w *Wallet) Balance() (amount float32) {
	return w.balance
}

func (w *Wallet) Withdraw(amount float32) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
