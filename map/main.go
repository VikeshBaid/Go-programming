package main

import (
	"fmt"
)

func main() {
	// map store data in key:value pair
	// map[string]int is a type. "[string]" shows the key and its type.
	// int show the type of value
	m := map[string]int{
		"james":      007,
		"moneypenny": 23,
	}

	fmt.Println(m)

	// to print the value for a key
	fmt.Println(m["james"])

	// to check whether the value is stored in map wth a key
	// "comma ok" idiom or ok is of bool type which returns false if key is not present in map
	v, ok := m["barnas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["james"]; ok{
		fmt.Println("this is the print statement inside if", v)
	}
	// add element to map

	m["todd"] = 34

	// for-range over map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// to delete a value from a map
	// delete(map, "key")

	delete(m, "james")
	fmt.Println(m)

	// map doesn't give any error if the key: value you are deleteing exists in map or not
	// so to check or lock-up it use "comma ok" idiom

	if v, ok := m["moneypenny"]; ok {
		fmt.Println("Value: ", v)
		delete(m, "moneypenny")
	}
	fmt.Println(m)
}
