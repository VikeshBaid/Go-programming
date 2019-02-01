package main

import (
	"fmt"
)

func main() {

	// assigning funciton to a variable is know as func expression
	f := func() {
		fmt.Println("my func expression")
	}
	f()

	g := func(n int) {
		fmt.Println("George Orwell", n)
	}
	g(1984)
}
