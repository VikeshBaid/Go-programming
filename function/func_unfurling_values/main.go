package main

import (
	"fmt"
)

func main() {
	// we can pass zero or more values to sum()
	// variadic means 0 or more values
	xi := []int{2, 3, 4, 5, 6, 8, 7, 9}
	x := sum(xi...)
	fmt.Println("the sum is, ", x)
}
// ...int represents 0 or more values pass to function
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
