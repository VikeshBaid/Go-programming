package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
	fmt.Println("main func")
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
