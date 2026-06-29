package main

import "testing"

// TODO change wants to notwant and get tests to pass
func TestName(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		got := Name("")
		want := any
		showCorrectMessage(t, got, want)
	})
	t.Run("'Paul' not allowed", func(t *testing.T) {
		got := Name("Paul")
		notWant := anything
		showCorrectMessage(t, got, want)
	})
}

func showCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
