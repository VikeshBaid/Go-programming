package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 72}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:4])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
