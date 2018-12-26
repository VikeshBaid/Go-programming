package main

import (
	"fmt"
)

// DECLARE the variable "y"
// ASSIGN the value 43
// declare + assign = initilization
var y = 43

// DECLARE there is a variable with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of the TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
// scope of z, a, b is all over the package as it is declared outside the func main
var z int

//var a float32
var b string = `James said,

"shaken,

not stirred` // back quotes " ` " is use to write the string laterals, one can use them to write multi-line string

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42 // scope of x is from here to end of the func main
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
	//fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	fmt.Println("coming from foo:", y)
}
