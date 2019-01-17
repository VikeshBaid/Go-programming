package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "bond",
		age:   25,
	}

	fmt.Println(p1)
}
