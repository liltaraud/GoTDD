package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(15)
		err := wallet.Withdraw(Bitcoin(7))

		assertBalance(t, wallet, Bitcoin(8))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		startBalance := Bitcoin(10)
		wallet.Deposit(startBalance)
		err := wallet.Withdraw(20)

		assertBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if got != want {
		t.Errorf("Wrong error message :\ngot \n\"%s\"\nwant \n\"%s\"", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("Unexpected error: %s", got.Error())
	}
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
