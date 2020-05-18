package main

import (
	"fmt"
	"os"
)

func main() {
	//Scope of the numchar is in the if statements
	if numchar, err := fmt.Printf("Hello!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n", numchar)
	}

	//here scope ends
}
