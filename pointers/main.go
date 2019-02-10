package main

import (
	"fmt"
)

func main() {
	a := 23
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	// below statement will show us the type of &a i.e. *int(pointer to an int)
	// it is completely different from type of a i.e. is int
	fmt.Printf("%T\n", &a)  // & gives you the address

	//var b int := &a // will give error as &a is of type *int and b is int
	// we can do the above statement as
	var b *int = &a
	fmt.Println(b)

	// to access the value stored at an address use *
	// also known as de-referencing
	fmt.Println(*b) // * gives you the value stored at an address when you have the address

	// we can also do like this
	fmt.Println(*&a)

	//changing the value using the address

	*b = 43
	fmt.Println(a)
	// above statement takes the address that is stored in variable b
	// and then *b will shows the value and then change it to the 
	// current assigned value
	// just like changeing the value of a variable by initializing it
	// with different value
}
