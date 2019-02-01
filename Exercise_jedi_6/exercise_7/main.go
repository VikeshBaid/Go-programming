package main

import (
	"fmt"
)

var g func([]int)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	s := func(xi []int) {
		total := 0
		for _, v := range xi {
			total += v
		}
		fmt.Println(total)
	}
	g = s
	g(ii)
	//fmt.Println(g)
	fmt.Printf("%T\n", g)
}
