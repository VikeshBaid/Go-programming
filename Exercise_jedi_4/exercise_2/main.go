package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 7, 84, 17}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", s)
}
