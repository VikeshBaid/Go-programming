package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 5, 54} // composite literal to declare slice
	fmt.Println(len(x))        // printing len of slice
	fmt.Println(x[0])          // printing value using index

	// printing value using for-range
	for i, v := range x {
		fmt.Println(i, v)
	}
}
