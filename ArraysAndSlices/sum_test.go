package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		// Slice
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got: %d, want: %d, given: %v", got, want, numbers)
		}

	})
	t.Run("Collection of any size", func(t *testing.T) {
		//Slice
		numbers := []int{4, 5, 1}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got: %d, want: %d, given: %v", got, want, numbers)
		}

	})

}
