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

	delete(mp, "no_dr")

	for k, v := range mp {
		fmt.Println(k, v)
	}
}
