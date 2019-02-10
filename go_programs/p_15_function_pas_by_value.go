// here we see that in Golang all the parameters are passed by values
// i.e it sends the copy of the value where ever it is needed
// so we can say that in Go , all parameters passed to a funciton are done so
// by VALUE

package main

import (
	"fmt"
	"math"
)

func dbl(val float64) {
	val = 2 * val
	fmt.Printf("in dbl function val change to %.5f\n", val)
}

func main() {
	p := math.Pi
	fmt.Printf("val of p before dbl funciton call %.5f\n", p)
	fmt.Println()
	dbl(p)
	fmt.Println()
	fmt.Printf("val of p after dbl dunciton call %.5f\n", p)
}
