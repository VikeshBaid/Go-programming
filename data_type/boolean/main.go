package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println("will give zero value(compiler assign value)", x)
	x = true
	fmt.Println(x)
}
