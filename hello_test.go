package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() tells the suite that this is a helper method
		// so the reported failure will be in the function call
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("you handsome delight")
		want := "Hello, you handsome delight"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

// 1. Tests need to be in a file with a name like blah_test.go
// 2. Test function must start with the word Test
// 3. Test function takes one argument-- t *testing.T

// t of type *testing.T is my hook into the framework, and I can call methods from t.
