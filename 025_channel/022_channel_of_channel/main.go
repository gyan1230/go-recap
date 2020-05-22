package main

import (
	"fmt"
	"time"
)

func emit(chanofchan chan chan string) {
	wordchan := make(chan string)
	chanofchan <- wordchan

	defer close(wordchan)

	words := []string{"hello", "world", "Hi", "How", "Are", "You"}
	t := time.NewTimer(100 * time.Microsecond)
	i := 0
	for {
		select {
		case <-t.C:
			return
		case wordchan <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}

		}
	}
}

func main() {
	channelch := make(chan chan string)
	go emit(channelch)

	wordch := <-channelch

	for word := range wordch {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
