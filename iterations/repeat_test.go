package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("it should return empty character if empty character is pass as argument", func(t *testing.T) {
		repeated := Repeat("", 5)
		expected := ""

		assertCorrectMessage(t, repeated, expected)
	})
}

// Run: go test -bench=.
// 136 ns/op means is our function takes on average 136 nanoseconds to run
// Which is pretttty ok! To test it ran it 10000000 times.

// by default Benchmarks are run sequentially.

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, tthe line number reported will be in our function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
