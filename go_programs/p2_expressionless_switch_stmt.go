package main

import (
	"fmt"
	"strings"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"USD", "US Dollar", "United States", 840},
}

func find(name string) {
	for i := 0; i < 10; i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found string", c)
		}
	}
}

func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found number", curr)
		}
	}
}

func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

func assertEuro(c Curr) bool {
	// below is a switch initializer which is used
	// to declare and initialize variables local to switch code block
	switch name, curr := "Euro", "EUR"; {
	case c.Name == name:
		return true
	case c.Currency == curr:
		return true
	}
	return false
}

func main() {
	find("Dinar")
	findAny(404)
}
