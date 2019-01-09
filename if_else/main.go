package main

import (
	"fmt"
)

func main() {
	// if with predefined boolean type
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	// if with expression
	if 2 == 2 {
		fmt.Println("004")
	}

	if 2 != 2 {
		fmt.Println("005")
	}

}
