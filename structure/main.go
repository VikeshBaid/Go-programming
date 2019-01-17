package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	age    int
	salary float64
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		age:    34,
		salary: 50000,
	}

	p2 := person{
		first:  "Miss",
		last:   "moneypenny",
		age:    32,
		salary: 54524,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.last, p1.first)
	fmt.Println(p2.age, p2.salary)
}
