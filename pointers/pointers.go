package pointers

import "errors"

type Wallet struct {
	balance float32
}

func (w *Wallet) Deposit(amount float32) {
	w.balance += amount
}

func (w *Wallet) Balance() (amount float32) {
	return w.balance
}

func (w *Wallet) Withdraw(amount float32) error {
	if amount > w.balance {
		return errors.New("not enough founds to withraw the amount")
	}
	w.balance -= amount
	return nil
}
