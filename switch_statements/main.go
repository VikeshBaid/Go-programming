package main

import (
	"fmt"
)

func main() {
	// if there is missing switch expression then it is considered as true by default which lead to evaluate only cases which evaluates to true. like below
	switch {
	case (1 == 0):
		fmt.Println("it will not print1")
	case true:
		fmt.Println("it will print")
		// here switch statements ends as the condition is true in case
		// if want to continue to next case we should use "fallthrough"
		fallthrough
	case false:
		fmt.Println("it will not print")
		// since fallthrough is not used it will end here 
		// no matter if the case is true or false as the first case 
		// is true
		fallthrough
	case (2 == 3):
		fmt.Println("it will print instead of  expression evaluates to false because of fall through in previous statement")
		fallthrough
	case (3 == 4):
		fmt.Println("it will print instead of false, as fallthrough is use in previous case")
		fallthrough
	case (4 == 4):
		fmt.Println("it will print since condition is true and fallthrough is use in previous case")
	default:
		fmt.Println("will not print since the last case is true and fallthrough is not used. If no case above is evaluated then only defalut statement is executed either there is fallthrough in previous statement or not")
	}
}
