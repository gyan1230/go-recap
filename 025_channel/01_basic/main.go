package main

import (
	"fmt"
	"time"
)

func main() {
	wordsch := make(chan string)
	go emit(wordsch)
	go emit(wordsch)
	go emit(wordsch)

	for v := range wordsch {
		fmt.Printf("%s\n", v)
	}

}

func emit(c chan string) {
	words := []string{"hello", "world", "Hi", "How", "Are", "You"}
	for _, v := range words {
		c <- v
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}
