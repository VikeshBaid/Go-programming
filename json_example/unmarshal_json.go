package main

import (
	"encoding/json"
	"fmt"
)

// `json:"field name"` is a tag we added so that unmarshal can say which is a
// field name to keep in Golang struct

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"james","Last":"bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v)
	}
}
