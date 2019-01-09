package main

import (
        "fmt"
)

func main() {
        for i := 1; i <= 5; i++ {
		// print space
		for j := 1; j < i; j++ {
			fmt.Printf(" ")
		}
		// print pattern
                for j := 5; j >= i; j-- {
                        fmt.Printf("%d", i)
                }
		// print pattern
                for k := 4; k >= i ; k-- {
                        fmt.Printf("%d", i)
                }
		fmt.Printf("\n")
        }
}
