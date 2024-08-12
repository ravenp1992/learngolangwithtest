package main

import (
	"fmt"
	"github.com/ravenp1992/learngolangwithtest/hello"
)

func main() {
	fmt.Println(hello.Greet("Raven"))

	x := [4]string{"raven", "paragas", "iya", "elia"}
	y := x[:]

	z := make([]string, len(x))
	copy(z, x[:])

	y[1] = "kristine"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)

	a := make([]int, 1e6)    // slice "a" with len = 1 million
	b := a[:2]               // even though "b" len = 2, it point to the same underlying array "a" points to
	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
