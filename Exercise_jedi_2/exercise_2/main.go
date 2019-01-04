package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (40 <= 42)
	c := (40 >= 42)
	d := (40 != 40)
	e := (`a` < `b`)
	f := (`a` > `b`)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
