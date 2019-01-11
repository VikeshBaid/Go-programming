package main

import (
	"fmt"
)

func main() {
	// []string also known as slice of string
	s1 := []string{"james", "bond", "chocolate", "hazelnut"}
	s2 := []string{"beyomkesh", "bakshi", "vanila", "tea"}
	
	// this is a multidimensioanl slice
	// also called as slice of a slice of string
	// both dimension should have same type
	m1 := [][]string{s1, s2}
	fmt.Println(m1)
}
