package main

import (
	"fmt"
)

func main() {
	favSport := "Pool"
	switch favSport {
	case "Cricket":
		fmt.Println("Favourite sport is Cricket")
	case "Pool":
		fmt.Println("Favourite sport is Pool")
	default:
		fmt.Println("no fav sport")
	}
}
