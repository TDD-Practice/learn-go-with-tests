package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Can create a Wallet with starting balance 0", func(t *testing.T) {
		newWallet := Wallet{}
		if newWallet.balance != 0 {
			t.Errorf("Failed to create %#v", newWallet)
		}
	})
}
