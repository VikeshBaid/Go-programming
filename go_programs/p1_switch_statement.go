package main

import (
	"fmt"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 12},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"USD", "US Dollar", "United States", 840},
	Curr{"EUR", "Euro", "Italy", 978},
}

func isDollar(curr Curr) bool {
	var result bool
	switch curr {
	default:
		result = false
	case Curr{"AUD", "Asutralian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}

// fallthrough statement will take the execution to the next case block' firat statement
// it will carry the info with it
// like  firat case will be match from below func, then Go compiler will enter in the case block
// where it find a fallthrough stetement which will take the compiler to the case block of next case block
// and start compiling the block from statement 1
func isDollar2(curr Curr) bool {
	switch curr {
	// switch block
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		// case block
		fallthrough
	case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}:
		// second case block
		fmt.Println("second case block of func isDollar2")
		fallthrough
	case Curr{"USD", "US Dollar", "United States", 840}:
		// third case block
		return true
	default:
		return false
	}
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4], currencies[10]:
		return true
	default:
		return false
	}
}

func main() {
	curr := Curr{"EUR", "Euro", "Italy", 978}
	if isDollar(curr) {
		fmt.Printf("%+v is Dollar currency\n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v is Euro currency\n", curr)
	} else {
		fmt.Println("currency is not Dollar or Euro")
	}

	do1 := Curr{"AUD", "Australian Dollar", "Australia", 36}
	if isDollar2(do1) {
		fmt.Println("Dollar currency found:", do1)
	}
}
