package main

import (
	"fmt"
)

func half(num *float64) {
	// *num means dereferenceing by use of pointer
	// we can access the value using pointer
	fmt.Printf("call half(%f) \n", *num)
	*num = *num / 2
}

func main() {
	num := 2.14585
	fmt.Printf("before half() func %f \n", num)
	half(&num)
	fmt.Printf("after half() func %f \n", num)
}
