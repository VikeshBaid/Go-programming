package main

import (
	"fmt"
)

//untypes constant(not defined explicitly)

const (
	a = 42
	b = 42.43
	c = "kamws"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
