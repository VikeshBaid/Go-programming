package main

import (
	"fmt"
)

func avg(num ...float64) float64 {
	n := len(num)
	t := 0.0
	for _, v := range num {
		t += v
	}

	return t / float64(n)
}

func sum(num ...float64) float64 {
	t := 0.0
	for _, v := range num {
		t += v
	}
	return t
}

func main() {
	fmt.Printf("avg(1,2.5,6.4) = %.2f\n", avg(1, 2.5, 6.4))
	points := []float64{9, 4, 3.7, 7.1, 9.2, 10}
	fmt.Printf("sum(%v) = %.2f\n", points, sum(points...))
}
