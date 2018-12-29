package main

import (
	"fmt"
)

type idli int

var x idli

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
