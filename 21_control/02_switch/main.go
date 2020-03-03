package main

import (
	"fmt"
	"os"
)

func main() {
	get := os.Args[1:]
	if len(get) != 1 {
		fmt.Println("enter some value")
		return
	}
	switch get[0] {
	case "polo", "vento", "t-roc":
		fmt.Println("Brand is Volkeswagen")
	case "Brezza", "Alto", "Swift", "Baleno":
		fmt.Println("Brand is Maruti")
	case "i10", "i20", "Creta":
		fmt.Println("Brand is Hyundai")
	default:
		fmt.Println("Brand is not available")
	}

}
