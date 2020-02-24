package main

import "fmt"

func main() {
	var m1 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	if res := valuecheck(m1, "two"); res {
		fmt.Println("key exist")
	} else {
		fmt.Println("no key")
	}
	fmt.Println(m1)
	delete(m1, "four")
	fmt.Println(m1)

}

func valuecheck(m1 map[string]int, key string) bool {
	_, ok := m1[key]
	return ok
}
