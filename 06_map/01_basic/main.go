package main

import (
	"fmt"
)

func main() {
	var m2 map[string]string
	m1 := make(map[int]string)
	fmt.Println(m1)
	m1[1] = "one"
	m1[2] = "two"
	fmt.Println(m1)
	m1[3] = "three"

	fmt.Println(m2)
	fmt.Println(m2 == nil)
	m2["gg"] = "dd"

	for i, v := range m1 {
		fmt.Println(i, v)
	}

}
