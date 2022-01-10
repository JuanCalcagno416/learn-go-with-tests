package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}
		w.Deposit(Bitcoin(10))

		want := Bitcoin(20)
		assertBalance(t, w, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		initialBalance := Bitcoin(100)
		w := Wallet{balance: initialBalance}

		want := Bitcoin(50)
		err := w.Withdraw(Bitcoin(50))
		assertNoError(t, err)
		assertBalance(t, w, want)

	})

	t.Run("Withdraw inssuficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, initialBalance)
	})

}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
