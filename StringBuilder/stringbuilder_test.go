package StringBuilder

import (
	"fmt"
	"testing"
)

func TestFullNameBuilder(t *testing.T) {
	got := fullNameBuilder("Hunter", "Weisenbach")
	expected := "Hunter Weisenbach"

	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}

func ExamplefullNameBuilder() {
	fmt.Println(fullNameBuilder("Hunter", "Weisenbach"))
	// Output: "Hunter Weisenbach"
}
