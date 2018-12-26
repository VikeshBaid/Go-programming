package main

import "fmt"

func main() {
	x := 42 // ':=' is a short hand declaration operator used to declare andassign with value simultaneously.
	fmt.Println(x)
	x = 99 // '=' is used instead of ':=' as variable x is already declared.
	y := 100 + 24
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
