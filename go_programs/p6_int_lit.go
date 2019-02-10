package main

import (
	"fmt"
)

func main() {
	vals := []int{
		1024,
		0x0FF1CE,
		0x8BADF00D,
		0xBEEF,
		0777,
	}

	for _, v := range vals {
		fmt.Printf("%d\n", v)
	}
}
