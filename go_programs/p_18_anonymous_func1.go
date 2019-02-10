package main

import (
	"fmt"
)

func main() {
	fmt.Printf("94 fahreniet = %.2f celcius \n", func(f float64) float64 {
		return (f - 32.0) * (5.0 / 9.0)
	}(94),
	)
}
