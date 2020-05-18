package main

import "fmt"

func main() {
	var nilmap map[string]int //adding keys to nil map error
	if nilmap == nil {
		fmt.Printf("Nil map..\n")
	}

	//Initialize map to add keys
	//one way to initialize
	var m1 = make(map[string]int)
	m1["jan"] = 31
	if m1 == nil {
		fmt.Printf("m1 Nil map..")
	} else {
		fmt.Println(m1)
	}

	//Initialize by using map literal
	var m2 = map[string]int{
		"feb": 28,
		"mar": 31,
	}
	if m2 == nil {
		fmt.Printf("m2 Nil map..")
	} else {
		fmt.Println(m2)
	}
	//  Create an empty map using map literal by leaving the curly braces empty
	m3 := map[string]int{} //functionally identical to using the make() function
	m3["mar"] = 31

	monthday := make(map[string]int)
	monthday["jan"] = 31
	monthday["feb"] = 29
	monthday["mar"] = 31
	monthday["apr"] = 30
	monthday["may"] = 31
	monthday["jun"] = 30
	monthday["jul"] = 31
	monthday["aug"] = 31
	monthday["sep"] = 30
	monthday["oct"] = 31
	monthday["nov"] = 30
	monthday["dec"] = 31

	//accessing value by key
	fmt.Printf("february no of days %d\n", monthday["feb"])
	//without check give zero value of the value of map
	// fmt.Printf("january no of days %d\n", monthday["january"])
	if day, found := monthday["january"]; !found {
		fmt.Printf("no month found\n")
	} else {
		fmt.Printf("january has %d days\n", day)
	}

	if day, found := monthday["jan"]; !found {
		fmt.Printf("no month found\n")
	} else {
		fmt.Printf("january has %d days\n", day)
	}
	//iterate over map
	for month, days := range monthday {
		fmt.Printf("%s month has %d days\n", month, days)
	}

}
