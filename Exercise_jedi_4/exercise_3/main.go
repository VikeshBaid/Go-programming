package main

import (
        "fmt"
)

func main() {
        s := []int{1, 2, 3, 4, 5, 6, 7, 7, 84, 17}
        fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}
