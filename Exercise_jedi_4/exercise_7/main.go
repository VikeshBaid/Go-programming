package main

import (
	"fmt"
)

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	m1 := [][]string{s1, s2}
	fmt.Println(m1)

	for i, s := range m1 {
		fmt.Println("record: ", i)
		for _, v := range s {
			fmt.Println("Values: ", v)
		}
	}
}
