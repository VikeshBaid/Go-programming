package main

import (
	"fmt"
)

func main() {
	var x [5]int // declaring an array x which has five elements and of type int.The length of array is part of it's type in array i.e [10]int and [20]int are different types in GO
	fmt.Println(x) // prints zero value of tyoe which is assigned
	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}
	fmt.Println(x)
}
