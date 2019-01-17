package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	favFlav []string
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		favFlav: []string{
			"chocolate",
			"martini",
		},
	}

	p2 := person{
		first: "miss",
		last:  "moneypenny",
		favFlav: []string{
			"strawberry",
			"rum",
		},
	}
	
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	//fmt.Println(m)

	for _, val := range m {
		fmt.Println(val.first)
		fmt.Println(val.last)
		for i, v := range val.favFlav {
			fmt.Println(i, v)
		}
		fmt.Println("---------")
	}
}
