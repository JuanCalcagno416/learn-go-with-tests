package main

import "testing"

func assertCorrectMessage(t testing.TB, res, expectedRes string) {
	t.Helper()
	if res != expectedRes {
		t.Errorf("res %q expected %q", res, expectedRes)
	}
}
func TestGreetingPrefix(t *testing.T) {
	var res, expectedRes string

	t.Run("When it receives Spanish as argument", func(t *testing.T) {
		res = GreetingPrefix("Spanish")
		expectedRes = "Hola, "
		assertCorrectMessage(t, res, expectedRes)
	})

	t.Run("When it receives English as argument", func(t *testing.T) {
		res = GreetingPrefix("English")
		expectedRes = "Hello, "
		assertCorrectMessage(t, res, expectedRes)
	})

	t.Run("When it receives French as argument", func(t *testing.T) {
		res = GreetingPrefix("French")
		expectedRes = "Bonjour, "
		assertCorrectMessage(t, res, expectedRes)
	})
}

func TestHello(t *testing.T) {
	var res, expectedRes string

	t.Run("saying hello to people in English", func(t *testing.T) {
		res = Hello("Nacho", "English")
		expectedRes = "Hello, Nacho"
		assertCorrectMessage(t, res, expectedRes)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		res = Hello("Nacho", "Spanish")
		expectedRes = "Hola, Nacho"
		assertCorrectMessage(t, res, expectedRes)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		res = Hello("Nacho", "French")
		expectedRes = "Bonjour, Nacho"
		assertCorrectMessage(t, res, expectedRes)
	})

	t.Run("say 'Hello, World' when there is no string provided", func(t *testing.T) {
		res = Hello("", "English")
		expectedRes = "Hello, World"
		assertCorrectMessage(t, res, expectedRes)

	})

}
