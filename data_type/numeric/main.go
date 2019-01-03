package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8 = -18  // int8 is valid from -128 to 127

func main() {
	//x := 42
	//y := 42.2315
	x = 7
	y = 7.05
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	// for runtime module read it from GoDoc.org/runtime
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
