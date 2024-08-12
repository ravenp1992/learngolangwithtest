package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestDiff(t *testing.T) {
	diff := Diff(4, 2)
	expected := 2

	if diff != expected {
		t.Errorf("expected '%d' but got '%d'", expected, diff)
	}
}

func ExampleDiff() {
	diff := Diff(5, 4)
	fmt.Println(diff)
	// Output: 1
}

// Testable Examples are compiled whenever tests are executed.
