package main

import "fmt"

func main() {
	//Declaring a map
	var m map[string]int
	if m == nil {
		fmt.Println("m is nil", m)
	}
	// Attempting to add keys to a nil map will result in a runtime error
	// m["one hundred"] = 100

	// Initializing a map using the built-in make() function
	var m2 = make(map[string]int)
	fmt.Println(m2)
	if m2 == nil {
		fmt.Println("m2 is nil")
	} else {
		fmt.Println("m2 is not nil")
	}
	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m2["hundred"] = 100
	fmt.Println(m2)

	// Initializing a map using a map literal
	var m3 = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m3)
	// create an empty map using a map literal by leaving the curly braces empty
	var m4 = map[string]int{}
	m4 = m3
	fmt.Println(m4)

	// Adding items (key-value pairs) to a map
	m4["three"] = 3
	m4["four"] = 4
	fmt.Println(m4)

}
