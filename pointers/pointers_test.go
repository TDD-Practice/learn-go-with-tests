package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Can create a Wallet with starting balance 0", func(t *testing.T) {
		newWallet := Wallet{}
		if newWallet.balance != 0 {
			t.Errorf("Failed to create %#v", newWallet)
		}
	})

	t.Run("Can deposit 100 and have a balance of 100", func(t *testing.T) {
		newWallet := Wallet{}

		newWallet.Deposit(100)

		if newWallet.balance != 100 {
			t.Errorf("Wallet %#v failed to deposit 100", newWallet)
		}
	})
}
