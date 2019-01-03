package main

import (
	"fmt"
)

const a int = 42
const b float64 = 43.42
const c string = "James Bond"

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
