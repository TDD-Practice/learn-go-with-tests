package pointers

type Wallet struct {
	balance float32
}

func (w *Wallet) Deposit(amount float32) {
	w.balance += amount
}

func (w *Wallet) Balance() (amount float32) {
	return w.balance
}
