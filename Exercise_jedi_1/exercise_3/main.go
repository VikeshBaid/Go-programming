package main

import (
	"fmt"
)

//package level access

var x = 42
var y = "James Bond"
var z = true

func main(){
	s := fmt.Sprintf("%v, %v, %v ", x, y, z)
	fmt.Println(s)
}
