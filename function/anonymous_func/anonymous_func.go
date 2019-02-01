package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("it is an anonymous function with no argument")
	}()

	func(n int) {
		fmt.Println("it is an anonymous function with one argument:")
		fmt.Printf("\t prints an integer %d\n", n)
	}(42)
}
