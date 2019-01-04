package main

import (
	"fmt"
)

// withing a constant declaration iota represents successive untyped integer constants. Value starts from 0

// iota is a predefined identifier which increments its value by 1 for every constant.

const (
	a = iota
	b
	c
	// can also write the above as 
	// a = iota
	// b = iota
	// c = iota

	// both will give same result
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
