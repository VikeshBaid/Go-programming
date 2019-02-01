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

//s is a object of type secretAgent and can access all the element of secretAgent
// person is struct type which is part of secretAgent, so the object of secretAgent can access the Identifiers of person struct
// speak function as a receiver i.e. secretAgent. it can access everything that comes under secreAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

// p is a object of struct type person. it can access the identifiers of struct person
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the Person speak")
}

// creating an interface
// the types which are associated with function speak()
// are also of type human.
// speak is of type interface which allows the other type value to be acces by human
type human interface {
	speak()
}

func bar(h human) {
	// using switch statement for assertion
	// here switch on with interface
	// here we are asseting, h is of this type
	switch h.(type) {
	case person:
		fmt.Println("person was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("secret agent was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "miss",
			last:  "moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println()
	fmt.Println(p1)
	p1.speak()
	
	// polymorphism
	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// conversion means we are converting a type into another type
	// but assertion means we know that the value is of particular type and saying that it is of particular type
}
