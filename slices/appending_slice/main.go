package main

import (
	"fmt"
)

func main() {
	x := []int{23, 456, 3532, 4, 5}
	// appending elements
	x = append(x, 43, 564, 43)
	fmt.Println(x)

	y := []int{43, 56, 2342}
	// appending slice to another
	x = append(x, y...) // "..." means take all the elemnts of y
	fmt.Println(x)

	// deleting elements from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
