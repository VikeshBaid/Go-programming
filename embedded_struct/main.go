package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type employee struct {
	person
	salary float64
}

func main() {
	em1 := employee{
		person : person{
			first: "James",
			last:  "bond",
			age:   32},

		salary: 52000,
	}

	em2 := employee{
		person : person{
			first: "miss",
			last: "grainger",
			age: 23,
		},
		salary: 0,
	}


	fmt.Println(em1)
	fmt.Println(em2)
}
