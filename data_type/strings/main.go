package main

import (
	"fmt"
)

func main() {
	s := "hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println("")
	//fmt.Printf("%q\n", s)
	//fmt.Printf("%x\n", s)
	//fmt.Println("")

	bs := []byte(s)  // this is slice of byte means a part of string
	fmt.Println(bs)
	fmt.Println("")
	fmt.Printf("%T\n", bs)
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%x ", s[i])
	//}

	//fmt.Println("")
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%d ", s[i])
	//}

	//fmt.Println("")
	//for i, v := range s {
	//	fmt.Printf("%#U \t %d", v, i)
	//}
}
