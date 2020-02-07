package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var books [4]string
	fmt.Printf("Size of books %d, values are %v\n", unsafe.Sizeof(books), books)

	books[0] = "go"
	books[1] = "ruby"
	books[3] = "python"
	fmt.Printf("Size of books %d, values are %v\n", unsafe.Sizeof(books), books)
	for i, v := range books {
		fmt.Println(i, ":", v)
	}
	type gyan [2]int
	var g gyan
	fmt.Printf("type of g is %T\n", g)
	fmt.Println(g)

}
