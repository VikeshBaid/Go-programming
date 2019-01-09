package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x)) // this is the length of slice allocated while declaration
	fmt.Println(cap(x)) // this gives the total capacity fo underlying array

	x[0] = 22
	x[9] = 32

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 5435)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// if extend the length of slice more than the size of unerlying array
	// complier delete the previous underlying array adn create a new one
	// with all the values present in previous array

	x = append(x, 352, 43)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
