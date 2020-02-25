package main

import "fmt"

type employee struct {
	fName, lName string
	age          int
}

func main() {
	//Just like other data types, you can declare a variable of a struct type like this -
	var e employee
	fmt.Println(e) // All the struct fields are initialized with their zero value

	//You can initialize a variable of a struct type using a struct literal like so -
	var e2 = employee{"gyan", "cchand", 32}
	fmt.Println(e2)

	// var e3 = employee{"gyan"}
	// fmt.Println(e3) //.\main.go:19:20: too few values in employee literal

	//Go also supports the name: value syntax for initializing a struct (the order of fields is irrelevant when using this syntax).
	var e4 = employee{
		fName: "jai",
		// lName: "pawar",	// All the uninitialized fields are set to their corresponding zero value -
		age: 28,
	}
	fmt.Println(e4)

}
