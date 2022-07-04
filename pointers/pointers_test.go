package pointers

import "testing"

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected an error to be returned")
	}

	if got != want {
		t.Errorf("got %q, but instead wanted %q", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Can create a Wallet with starting balance 0", func(t *testing.T) {
		newWallet := Wallet{}
		if newWallet.balance != 0 {
			t.Errorf("Failed to create %#v", newWallet)
		}
	})

	t.Run("Can deposit 100 and have a balance of 100", func(t *testing.T) {
		newWallet := Wallet{}
		newWallet.Deposit(100.00)
		got := newWallet.Balance()
		want := float32(100.00)
		if got != want {
			t.Errorf("Wallet %#v failed to deposit 100. Deposit is %.2f", newWallet, newWallet.Balance())
		}
	})
}
func TestTableWallet(t *testing.T) {
	testCases := []struct {
		desc   string
		w      Wallet
		amount float32
		want   float32
	}{
		{
			desc:   "Can deposit 100 and have a balance of 100",
			w:      Wallet{},
			amount: float32(100),
			want:   float32(100.00),
		},
		{
			desc:   "Can deposit 10 and have a balance of 10",
			w:      Wallet{},
			amount: float32(100),
			want:   float32(100.00),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.w.Deposit(tC.amount)
			got := tC.w.Balance()
			if got != tC.want {
				t.Errorf("Wallet %#v failed to deposit 100. Deposit is %.2f", tC.w, tC.w.Balance())
			}

			tC.w.Withdraw(tC.amount)
			got = tC.w.Balance()
			if got != 0 {
				t.Errorf("Wallet %#v failed to withdraw 100. Deposit is %.2f", tC.w, tC.w.Balance())
			}
		})
	}
}

func TestOverdrafProtection(t *testing.T) {

	t.Run("Can withdraw 10 form 100 more than Balance()", func(t *testing.T) {
		newWallet := Wallet{15}
		err := newWallet.Withdraw(10)
		if err != nil {
			t.Errorf("Should be able to withdraw form sufficient funds in wallet %#v", newWallet)
		}
	})

	t.Run("Cannot withdraw more than Balance()", func(t *testing.T) {
		newWallet := Wallet{15}
		err := newWallet.Withdraw(newWallet.Balance() + 1)
		assertError(t, err, ErrInsufficientFunds)
	})
}
