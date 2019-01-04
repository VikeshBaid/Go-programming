package main

import (
	"fmt"
)

// using iota we do the same thing as in func main
const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)

	// we can say that iota increase the value by ncrementing 1 bit i.e
	// 0 = 0000 0000
	// 1 = 0000 0001
	// 2 = 0000 0010
	// 3 = 0000 0011
	// 4 = 0000 0100

	//from above we can say that iota increases the value in bits and give it's value as decimals

)

func main() {
	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
