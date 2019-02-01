package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	x := evensum(ii)
	fmt.Println(x())

}

func evensum(yi []int) func() int {
	var ei []int

	for _, v := range yi {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return func() int {
		total := 0
		for _, v := range ei {
			total += v
		}
		return total
	}
}
