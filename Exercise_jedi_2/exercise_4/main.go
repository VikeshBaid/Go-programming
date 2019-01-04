package main

import (
	"fmt"
)

func main() {
	s := 98
	fmt.Printf("%d\t\t%b\t\t%#x\n", s, s, s)

	bw := s << 1
	fmt.Printf("%d\t\t%b\t%#x\n", bw, bw, bw)
}
