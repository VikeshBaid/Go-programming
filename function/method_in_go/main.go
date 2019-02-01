package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// here "s secreAgent" is a receiver.
// it works just like an argument of a function.
// By the below definition speak method can be called by any value of type secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)

	// calling METHOD speak() by VALUE sa1 of TYPE secretAgent
	sa1.speak()
}
