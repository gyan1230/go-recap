package main

import "fmt"

func main() {
	const day int = 12
	var i int // auto initilize to zero value
	fmt.Println(i)

	// const ok bool	// error: missing value, need to initialize
	const pi float32 = 3.14

	// a, b := 5, 0
	// fmt.Println(a / b) //runtime error can't detect

	const a1, b1 int = 5, 0
	// fmt.Println(a1 / b1) // detect error

	//multiple constant declaration
	const a2, b2 = 1, 2
	const (
		a3 = 123
		b3
		c3
	)
	fmt.Println(a3, b3, c3)
}
