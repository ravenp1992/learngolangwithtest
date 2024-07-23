package main

import "fmt"

// Constants
const myConstX int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

func main() {
	// Integer literals
	base16 := 0x01
	base10 := 8
	base8 := 0o07
	base2 := 0b1000
	fmt.Printf("%v\n", base16)
	fmt.Printf("%v\n", base10)
	fmt.Printf("%v\n", base8)
	fmt.Printf("%v\n", base2)

	// Floating point literal
	fmt.Printf("%v\n", 6.03e23)

	// Rune literal
	myRune := 'R'
	fmt.Println(myRune)

	// String literals
	myString := "Raven"
	myRawString := `Raven "Paragas"`
	fmt.Println(myString)
	fmt.Printf("%s\n", myRawString)

	// Booleans
	var flag bool // no value assigned, set to false
	var isAwesome = true
	fmt.Println(flag)
	fmt.Println(isAwesome)

	// Dividing a nonzero floating point variable by 0 returnn +Inf or -Inf
	var f1 float64 = 1.5
	var f2 float64 = -1.5
	fmt.Println(f1 / 0)
	fmt.Println(f2 / 0)

	// Dividing a floating point variable set to 0 by 0 return NaN (Not a Number)
	var myFloat float64 = 0
	fmt.Println(myFloat / 0)

	{
		var myString string
		fmt.Printf("%q\n", myString)

		// Strings in Go are immutable, you can reassign the value of a string variable, but you cannot change the value of the string that is assgined to it.
		// myString[0] = "R"

		myString = "Hello"
		fmt.Println(myString[1:])
	}

	{
		// Explicit Type Conversion
		var x int = 10
		var y float64 = 30.2
		var z float64 = float64(x) + y
		var d int = x + int(y)
		fmt.Printf("%f\n", z)
		fmt.Printf("%d\n", d)
	}

	{
		// var versus :=

		// var x int = 10 // The most verbose wayt to declare a variable in Go.
		// var x = 10 // Since the default type of an integer literal is int, the following declares x to be a variable of type int.
		// var x int // will be 0 because its the zero value of any numeric
		// var x, y int = 10, 20 // Multiple variables at once with var, and they can be of the same type:
		// var x, y int // all zero value of the same type.
		// var x, y = 10, "hello" // or of different types.
		// fmt.Println(x, y)

		// There's one more way to use var. If you are declaring multiple variables at once, you can wrap them in a declaration list:
		// var (
		// 	x    int
		// 	y        = 20
		// 	z    int = 30
		// 	d, e     = 40, "hello"
		// 	f, g string
		// )
		// fmt.Println(x, y, z, d, e, f, g)

		// Go also supports a short declaration format.

		// The following two statements do exaclty the same thing:
		// var x = 10
		// x :=

		// var x, y = 10, "hello"
		// x, y := 10, "hello"

		x := 10
		fmt.Println(x)
		x, y := 30, "hello"
		fmt.Println(x, y)

	}

	{

		const z = 20 * 10

		const y = "hello"

		fmt.Println(myConstX)
		fmt.Println(y)

		// error
		// myConstX = myConstX + 1
		// y = "bye"
	}

	{
		// Unused Variables
		x := 10
		x = 20
		fmt.Println(x)
		x = 30
	}
}
