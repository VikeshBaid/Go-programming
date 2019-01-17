package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	favFlav    []string
}

func main() {
	p1 := person{
		first_name: "James",
		last_name:  "Bond",
		favFlav:    []string{"chocolate chip", "martini"},
	}

	p2 := person{
		first_name: "Shrink",
		last_name:  "Gupta",
		favFlav:    []string{"strawberry", "vanilla"},
	}

	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	fmt.Printf("Favorite Flavour:\n")
	for _, v := range p1.favFlav {
		fmt.Printf("\t\t%s\n", v)
	}

	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	fmt.Printf("Favorite Flavour:\n")
	for _, v := range p2.favFlav {
		fmt.Printf("\t\t%s\n",  v)
	}
}
