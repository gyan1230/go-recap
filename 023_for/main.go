package main

import "fmt"

func main() {
	i := 0
	//One way
	for {
		i++
		if i > 5 {
			break
		}
		fmt.Printf("Hello\n")
	}

	//second
	i = 0
	for i < 5 {
		fmt.Printf("hello\n")
		i++
	}

	//third
	for counter := 0; counter < 5; counter++ {
		fmt.Printf("world\n")
	}
}
