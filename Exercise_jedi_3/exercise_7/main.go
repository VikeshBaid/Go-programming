package main

import (
	"fmt"
)

func main() {
	x := 23

	if x == 40 {
		fmt.Println("inside if")
	} else if x == 23 { 
		fmt.Println("inside else if")
	} else {
		fmt.Println("inside else")
	}
}
