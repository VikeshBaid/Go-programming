package main

import (
	"fmt"
)

func main() {
	// x as a function in it which can be run by x()
	//x := bar()
	//fmt.Printf("%T\n", x)

	// i as the output (which is a integer) of x in it
	//i := x()
	//fmt.Println(i)

	// above things ca also be done by following ways
	//first way
	// x:= bar()
	// fmt.Println(x())	// to get the output

	// second way
	fmt.Println(bar()()) // printing output without using variable

}

// func() int is a return type for the function bar
// bar() is returning a function
func bar() func() int {
	return func() int {
		return 451
	}
}
