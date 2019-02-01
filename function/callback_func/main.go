// calling back a function from a function
// using a functin as an argument to another function

package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...) // unfurl the values of ii to pass the value in function
	fmt.Println("sum of numbers:", s)

	s2 := evensum(sum, ii...) // unfurl the values of ii to pass the value in function
	fmt.Println("sum of even numbers:", s2)

	s3 := oddsum(sum , ii...)// unfurl the values of ii to pass the value in function
	fmt.Println("sum of odd numbers:", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func evensum(f func(ai ...int) int, y ...int) int {
	var vi []int
	for _, v := range y {
		if v%2 == 0 {
			vi = append(vi, v)
		}
	}
	return f(vi...)
}

func oddsum(f func(bi ...int) int, z ...int) int {
	var oi []int
	for _, v := range z {
		if v % 2 != 0 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}
