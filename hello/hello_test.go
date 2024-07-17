package hello

import "testing"

// Writing tests

// It need to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test fucntion takes one argument only t *testing.Tt
// To use the *testing.T type, you need to import "testing", like w did with fmt in the other file

// Discipline
// - Write a test
// - Make the compiler pass
// - Run the test, see that it fails and check the error message is meaningful
// - Write enough code to make the test pass
// - Refactor

func TestHello(t *testing.T) {
	t.Run("say 'Hello, World' when an emtpy string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Raven", "")
		want := "Hello, Raven"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish to people", func(t *testing.T) {
		got := Hello("Raven", "Spanish")
		want := "Hola, Raven"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french to people", func(t *testing.T) {
		got := Hello("Raven", "French")
		want := "Bonjour, Raven"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in tagalog to people", func(t *testing.T) {
		got := Hello("Raven", "Tagalog")
		want := "Kumusta, Raven"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, tthe line number reported will be in our function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
