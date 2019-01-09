package main

import (
	"fmt"
)

func main() {
	i := 1993
	for {
		if i > 2019 {
			break
		}
		fmt.Println(i)
		i++
	}
}
