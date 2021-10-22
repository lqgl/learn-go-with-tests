package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(1000)}

		wallet.Withdraw(100)

		got := wallet.Balance()
		want := Bitcoin(900)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})
}
