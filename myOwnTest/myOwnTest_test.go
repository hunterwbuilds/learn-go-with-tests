package main

import "testing"

// TODO change wants to notwant and get tests to pass
func TestName(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		got := Name("")
		notWant := ""
		showCorrectMessage(t, got, notWant)
	})
	t.Run("'Paul' not allowed", func(t *testing.T) {
		got := Name("Paul")
		notWant := "Paul"
		showCorrectMessage(t, got, notWant)
	})
}

func showCorrectMessage(t testing.TB, got, notWant string) {
	t.Helper()
	if got == notWant {
		t.Errorf("got: %q, want: %q", got, notWant)
	}
}
