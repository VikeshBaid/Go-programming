package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Hello Foo")
}

func bar() {
	fmt.Println("Hello Bar")
}
