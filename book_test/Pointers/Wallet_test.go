package pointers

import "testing"

func TestWalll(t *testing.T) {
	wallet := Wallet{}

	t.Run("add money", func(t *testing.T) {

		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw money", func(t *testing.T) {

		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		checkBalance(t, wallet, 0)
	})

	t.Run("withdrwa essuviffund", func(t *testing.T) {
		err := wallet.Withdraw(Bitcoin(0))

		assertError(t, err, ErrInsufficientFunds.Error())
		checkBalance(t, wallet, 0)

	})
}
func checkBalance(t testing.TB, wallet Wallet, want Bitcoin) {

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want on")
	}
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()

	//we want error if it's not present below code shouldn't execute
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err.Error() == want {
		t.Error(err.Error())
	}
}
