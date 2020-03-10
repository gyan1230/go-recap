package main

import "fmt"

func main() {
	var sum int
	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Afer Continue")
		sum += i
		if sum > 30 {
			break
		}
		fmt.Println("Afer break")
		fmt.Println(sum)
	}
	fmt.Println(sum)

	for x := 0; x <= 5; x++ {
		for y := 0; y < 5; y++ {
			if x == 2 && y == 3 { //skip printing "inner loop when x is 2 and y is 3"
				continue
			}
			fmt.Println("inner loop X:", x, " Y:", y)
			sum = x + y
		}
		if x == 3 {
			break
		}
		fmt.Println("outer loop X: ", x)
	}
}
