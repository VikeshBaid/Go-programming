package main

import (
	"fmt"
)

var y string
var z int

func main() {
	// DECLAE a VARIABLE o be of certain TYPE
	// and then ASSIGN a VALUE of that type to variable

	fmt.Println("printing the value of y", y, "ending")

	//  %T	a Go-syntax representation of the type of the value
	// and "\n" is a new line character
	fmt.Printf("%T\n", y)

	y = "Bond, James"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z) // return zero value of z as it is declared with type butnot assign any value
	fmt.Printf("%T\n", z)
}
