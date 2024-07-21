package arrayandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

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
	t.Run("Returning a new slice containing the totals for each slice passed in.", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// Go does not let you use equality operators with slices.
		// if got != want {
		// 	t.Errorf("got %v want %v", got, want)
		// }

		// if !slices.Equal(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	// checkSums := func(t testing.TB, got, want []int) {
	// 	t.Helper()
	// 	if !slices.Equal(got, want) {
	// 		t.Errorf("got %v want %v", got, want)
	// 	}
	// }

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// func BenchmarkSum(b *testing.B) {
// 	// for i := 0; i < b.N; i++ {
// 	Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// 	// }
// }
