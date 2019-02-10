package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	x := evensum(ii)
	fmt.Println(x())

	y := evensum1()
	yy := y(ii)

	fmt.Println(yy)

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

func evensum1() func(xi []int) int {
	return func(xi []int) int {
		total := 0
		for _, v := range xi {
			if v%2 == 0 {
				total += v
			}
		}
		return total
	}
}
