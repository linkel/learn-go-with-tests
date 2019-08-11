package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("you handsome delight")
	want := "Hello, you handsome delight"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// 1. Tests need to be in a file with a name like blah_test.go
// 2. Test function must start with the word Test
// 3. Test function takes one argument-- t *testing.T

// t of type *testing.T is my hook into the framework, and I can call methods from t.
