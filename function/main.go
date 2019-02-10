package main

import (
	"fmt"
)

func main() {
	foo()
	boo("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flaming")
	fmt.Println(x)
	fmt.Println(y)
	ax := sum(2,3,4,5,7,8,8,9)
	fmt.Println("the sum is ,", &ax)
}

// func (r reciever) identifier(parameter(s)) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

// everthing in Go is PASS BY VALUE
// function with one parameter s
func boo(s string) {
	fmt.Println("hello,", s)
}

// function that is returning a value of type string
func woo(s string) string {
	return fmt.Sprint("hello from woo, ", s)
}

// function with multiple return(string and bool)
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn," ", ln, `, says "hello"`)
	b := true
	return a, b
}

// function with variadic parameters
// "..." operator shows multiple values and then type int is type of all values
// x is of type []int(slice of int)
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
