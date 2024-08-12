package arraysandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
		// if !slices.Equal(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 5})
		want := []int{0, 14}

		checkSums(t, got, want)
		// if !slices.Equal(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
	})
}

// go test -cover
