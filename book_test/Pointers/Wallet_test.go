package pointers

import "testing"

func TestWalll(t *testing.T) {
	wallet := Wallet{}

	checkBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("add money", func(t *testing.T) {

		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw money", func(t *testing.T) {

		wallet.Withdraw(Bitcoin(10))

		checkBalance(t, wallet, 0)
	})
}
