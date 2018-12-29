package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println("a =",a)
	fmt.Printf("type of a %T\n", a)
	fmt.Println("b = ", b)
	fmt.Printf("type of b %T(package.type)\n", b)

	// to assign a value of one type to another,you have to do conversion
	a = int(b) // type conversion
	fmt.Println("after type conversion and assigning the value of b to a", a)
	fmt.Printf("%T\n", a)

}
