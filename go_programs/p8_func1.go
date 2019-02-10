package main

import (
	"fmt"
)

func avogardo() float64 {
	return 6.02214129e23
}

func main() {
	fmt.Printf("avagardo() = %e 1/mol\n", avogardo())
}
