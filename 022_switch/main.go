package main

import (
	"fmt"
	"os"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog"
	vow := 0
	con := 0
	tee := 0
	spa := 0
	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vow++
		case ' ':
			spa++
		case 't':
			tee++
			fallthrough
		default:
			con++
		}
	}
	n, err := fmt.Printf("vowels: %d, consonants: %d, tee: %d, spaces: %d\n", vow, con, tee, spa)
	switch {
	case err != nil:
		os.Exit(1)
	case n > 0:
		fmt.Println("len of string is: ", n)
	default:
		fmt.Println("Ok!")
	}

}
