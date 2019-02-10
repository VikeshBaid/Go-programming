// when functions identifiers appear without parenthese, it is treated as a regular variable with a type and a value

package main

import (
	"fmt"
)

func add(op0 int, op1 int) int {
	return op0 + op1
}

func sub(op0 int, op1 int) int {
	return op0 - op1
}

func main() {
	// since add and sub is used without parentheses so it is treated
	// as a variable of with type(which in this case is given before)
	// and a value
	var opAdd func(int, int) int = add
	opSub := sub
	fmt.Printf("add(12,44) = %d\n", opAdd(12, 44))
	fmt.Printf("sub(99,13) = %d\n", opSub(99, 13))
	fmt.Printf("%T\n", sub)
}
