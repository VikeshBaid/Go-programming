package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		r := i%4
		fmt.Printf("%d gives reminder %d\n", i, r)
	}
}
