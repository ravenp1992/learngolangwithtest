# Promitive Types and Declarations

## Built-in Types

Using these types idiomatically is sometimes a challenge for developers who are transitioning from another language. We are going to look at these types and see how they work best in Go.

- booleans
- integers
- floats
- strings

## The Zero Value

Go, like most modern languages, assigns a default zero value to any variable that is declared but not assgined a value. Having an explicit zero value makes code clearer and removes a source of bugs found in C and C++ programs.

## Literals

A Literal in Go refers to writing out a number, character, or string. There are four common kinds of literals that you'll find in Go programs.

- **_Integer iterals_** are sequences of numbers; they are normally base 10, but different prefixes are used to indicate other bases: 0b for binary (base 2), 0o for octal(base 8), or 0x for hexadecimal (base 16)
- **_Floating point literals_** have decimal points to indicate the fractional portion of the value. They can also have an exponent specified with the letter e and a positive or negative number (such as 6.03e23)
- **_Rune literals_** represent characters and are surrounded by single quotes. Unlike many other languages, in Go single quotes and double quotes are _not_ interchangeable. Rune literals can be written as single Unicode character ('a'), 8-bit octal numbers ('\141'), 8-bit hexadecimal numbers ('\x61'), 16-bit hexadecimal numbers ('\u0061'), or 32-bit Unicode numbers ('\U00000061'). There are also several slash escaped run literals, with most useful ones being newline ('\n'), tab ('\t'), single Quote ('\\''), double quote ('\\"'), and backslash ('\\')
  )
- **_String literals_** Most of the time, you should use double quotes to create an interpreted string literal (e.g "Greetings and Salutations"). If you need to include backslashes, double quotes, or newlines in your string, use a _raw string literal_. These are delimitted with backquotes (\`) and can contain any literal character except backquotes. (e.g \`Greeting and "Salutation"\`).

## Booleans

The bool type represents Boolean variables. Variables of bool type can have one of two values: true or false. The zero value for a bool is false:

```go
var flag bool // no value assigned, set to false
var isAwesome = true
```

## Nmeric Types

Go has a large number of numer types: 12 different types (and a few special names) that are grouped into three categories.

#### Integer types

Go provdes both signed and unsigned integers in a variety of sizes, from one to four bytes.

| Type name | Value range                                 |
| --------- | ------------------------------------------- |
| int8      | -128 to 127                                 |
| int16     | -32768 to 32767                             |
| int32     | -2147483648 to 2147483647                   |
| int64     | –9223372036854775808 to 9223372036854775807 |
| uint8     | 0 to 255                                    |
| uint16    | 0 to 65536                                  |
| uint32    | 0 to 4294967295                             |
| uint64    | 0 to 18446744073709551615                   |

#### The special integer types

Go does have some special names for intege types. A byte is an alias for uin8; it is legal to assign, compare, or perform mathemetical operations between a byte and uint8. However, you rarely see uint8 used in Go code; just call it a byte.

The second special name is int. On 32-bit CPU, int is a 32-bit signed integer like int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64. Because int isn't consistent from platform to platform, it is a compile-time error to assigned, compare, or perform mathematical operations between an int and an int32 or int64 without type conversion. Integr literals defualt to being of int type

The Third special name is uint. It follows the same rules as int, only it is unsigned.

There are two other special names for integer types, rune and uintptr.

#### Choosing which integer to use

- If you are working with a binary file format or network protocol that has an integer of a specific size or sign, use the corresponding integer type.
- If you are writing a library function that should work with any integer type, write a pair of functions, one with int64 for the parameters and variables and the other with uint64

> The reason why int64 and uint64 are the idiomatic choice in this situation is that Go doesn't have generics (yett) and doesn't have function overloading. Without these feattures, you'd need to write many functions with slightly different names to implement your algorithm. Using int64 and uint64 means that you can write the code once and let your caller use type conversions to pass value in and convert data that's returned.

> You can see this pattern in the Go standard library with the functions FormatInt/FormatUint and ParseInt/ParseUint in the strconv package. There are other situations, like in the math/bits package, where the size of the integer matters. In those cases, you need to write a separate function for every integer type.

- In all other cases, just use int

#### Integer operators

Go integers support the usual arithmetic operators: +,-,\*,/, with % for modulus. The result of an integer division is an integer; if you want to get a floatting point result, you need to use a type conversion to make your integers into floating point numbers. Also be careful not to divide an integer by 0; this causes a panic.

> Integer division in Go follows truncation toward zero;

Yo ucan compare integers with == !=, >, >=, <, and <=, go also has abit-manipulation operators for integers. You can bit shift left and right with << and >>, or do bit masks with & (logical AND), | (logical OR), ^ (logical XOR), and &^ (logical AND NOT). You can also combine all of the logical operators with = to modify a variable: &=, |=, ^=, &^=, <<=, and >>=.

#### Floating point types

There are two floating point types in Go.

| Type name | Largest absolute value                          | Smallest (nonzero) absolute value              |
| --------- | ----------------------------------------------- | ---------------------------------------------- |
| float32   | 3.40282346638528859811704183484516925440e+38    | 1.401298464324817070923729583289916131280e-45  |
| float64   | 1.797693134862315708145274237317043567981e +308 | 4.940656458412465441765687928682213723651e-324 |

> Always use float64. Don't worry about the difference in memory size unless you have used the profiler to determine that it is a significant source of problems.

> IEEE 754\
> As mentioned earlier, Go (and most other programming languages) stores floating pointer numbers using a specification called IEEE 754.

> Many programmers learn at some point how integers are represented in binary (rightmost position is 1, next is 2, next is 4, and so on). Floating pointer numbers are very different. Out of the 64 bits above, one is used to represent the sign (positive or negative), 11 are used to represent a base two exponent, and 52 bits are used to represent the number in a normalized format (called the _mantissa_ )

#### Complex types (you're probably not going to use there)

## A tastet of Strings and Runes

This brings us to strings. Like most modern languages, Go includes strings as a built-in type. The zero value for a string is the empty strinng. Go supports Unicode;\
Like integers and floats, strings are compared for equality using ==, difference with !=, or ordering with >, >=, <, or <=. They are concatednated by using the + operator.

Go also has a type that represents a single code point. The _rune_ type is an alias for the int32 type, just like byte is an alias for uint8. As you could probably guess, a rune literal's default type is a rune, and a string literal's default type is a string.

> If you are referring to a character, use the rune type, not the in32 type. They might be tthe same to compiler, but you want to use the type that clarifies the intent of your code.

## Explicit Type Conversion

Go doesn't allow automatic type promotion between variables.

No another type can be converted to a bool, implicitly or explicity. If you want to convert another data type to boolean, you must use on of the comparison operators

## var versus :=

- If your are declaring a variable at package level, you must use var because := is not legal outside of functions.

There are some situations within functions where you should avoid :=;

- When initializing a variable to its zero value, use var x int. This makes it clear that the zero value is intended.
- When assigning an untyped constant or a literal to a variable and the default type for the constant or literal isn't the type you want for the variable, use the long var form with the type specified. While it is legal to use a type conversion to specify the type of the value and use := to write x := byte(20), it is idiomatic to write var x byte = 20.
- Because := allows you to assign to both new and existing variables, it sometimes creates new variable when you think you are reusing existing ones. In those situations, explicitly declare all of your new variables with var to make it clear which variables are new, and then use the assignment operator(=) to assign to both new and old variables.

> Avoid declaring variables outside of functions because they complicate data flow analysis.

## Using const

When developer learn a new programming language, they try tto map familiar concepts. Many languages have a way to declare a value is immutable. In Go, this is done with the const keyword. At first glance, it seems to work exactly like other languages.

However, const in Go is very limited. Constant in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time. This means that they can be assigned:

- Numeric literals
- true and false
- Strings
- Runes
- The built-in function complex, real, imag, len, and cap
- Expressions that consist of operator and the preceding values

Go doesn't provide a way to specify that a value calculated at runtime is immutable.

> Constants in Go are a way to give names to literals. There is no way in Go to declare that a variable is immutable.

## Typed and Untyped Constants

Constants can be typed or untyped. An untyped constant works exactly like a literal; it has no type of its own, but does have a default type that is used when no other type can be inferred. A type constant can only be directly assigned to a variable of that type.

```go
// Here's what an untyped constant declaration looks like:
const x = 10
// All of the following assignments are legal:
var y int = x
var z float64 = x
var d byte = x

// Here's what typed constant declaration looks like:
cons typedX int = 10
// This constant can only be assigned directly to an int. Assigning it to any other type produces a compile-time error like this:
// cannot use typedX (type int) as type float64 in assignment
```

## Unused Variables

Go requirements is that every declared local variable must be read. It is a compile-time error to declare a local variable and to not read its value.

The compiler's unused variable check is not exhaustive. As long as a variable is read once, the compiler won't complain, even if there are writes to the variable that are never read.

> The Go compiler won't stop you from creating unread package level variables. This is one more reason why you should avoid creating package-level variables.

> Unused Constants
> Perhaps surprisingly, the Go compiler allows you to create unread constants with const. This is because constants in Go are calculated at compile time and cannot have any side effects. This makes them easy to eliminate: if a constant isn't used, it is simply not included in the compiled binary.

## Naming Variables and Constants

- Go uses camel case (names like indexCounter or numberTies) when an identifier name consist of multiple words.
- Within a function, favor short variable names. The smaller the scope for a variable, the shorter the name that's used for it.
- When naming variables and constant in the package block, use more descriptive names. The type should still be excluded from the name, but since the scope is wider, you need a more complete name to make it clear what the value represents.
