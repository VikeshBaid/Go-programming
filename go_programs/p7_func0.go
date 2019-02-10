package main

import (
	"fmt"
	"math"
)

func printPi() {
	// funciton literal : pass a function as an argument in a method
	fmt.Printf("printPi() %v\n", math.Pi)
}

func main() {
	printPi()
}
