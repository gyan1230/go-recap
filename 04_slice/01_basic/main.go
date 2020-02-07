package main

import (
	"fmt"
	"strings"
)

func main() {
	var names []string
	var new []string
	// names[0] = "ggyyaann"	//can't get/set values
	names = append(names, "gyan", "jai", "vinod")
	fmt.Println(names)
	fmt.Println(strings.Repeat("*", 40))
	new = append(new, "1", "2", "3")
	names = append(names, new...)
	fmt.Println(names)

	games := make([]string, 0, 5)
	games = append(games, "mario", "contra", "pacman")
	for i, v := range games {
		fmt.Println(i, v)
		fmt.Printf("size is %d, capacity is %d\n", len(games), cap(games))
	}
}
