package main

import (
	"fmt"
)

func main() {
	mp := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(mp)
	fmt.Println()
	for k, v := range mp {
		fmt.Println("This is the record for ", k)
		for j, val := range v {
			fmt.Printf("\t %d \t  %v\n", j, val)
		}
	}
}
