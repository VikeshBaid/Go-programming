// anonymous struct
package main

import (
	"fmt"
)

func main() {
	s := struct {
		first    string
		last     string
		friends  map[string]int
		favDrink []string
	}{
		first: "James",
		last:  "Bond",
		friends: map[string]int{
			"M": 2254855,
			"Q": 2465813,
		},
		favDrink: []string{
			"martini",
			"rum",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.last)
	for k, val := range s.friends{
		fmt.Println(k, val)
	}

	for i, v := range s.favDrink {
		fmt.Println(i, v)
	}
}
