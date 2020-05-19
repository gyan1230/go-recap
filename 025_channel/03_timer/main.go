package main

import (
	"fmt"
	"time"
)

func main() {
	wordch := make(chan string)
	go emit(wordch)
	for w := range wordch {
		fmt.Printf("%s ", w)
	}
}

func emit(wordch chan string) {
	words := []string{"hello", "world", "Hi", "How", "Are", "You"}
	defer close(wordch)
	t := time.NewTimer(1 * time.Second)
	i := 0
	for {
		select {
		case wordch <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-t.C:
			return

		}
	}
}
