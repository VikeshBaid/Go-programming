package main

import (
	"fmt"
)

// has q and r is given in function signature
// no need to declare new variable inside function
func div(dvdn, dvsr int) (q, r int) {
	// r and q are taken from funciton signature
	// that's why no short declaration operator is used here
	// with respect to p_13
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return
}

func main() {
	q, r := div(71, 5)
	fmt.Println(q, r)
}
