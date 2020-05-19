package main

import "fmt"

func main() {

	wordschn := make(chan string)
	done := make(chan bool)

	go emit(wordschn, done)
	for i := 0; i < 12; i++ {
		fmt.Printf("%s\n", <-wordschn)
	}
	done <- true
}
func emit(wordch chan string, done chan bool) {
	words := []string{"hello", "world", "Hi", "How", "Are", "You"}
	i := 0
	for {
		select {
		case wordch <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			close(done)
			return
		}
	}

}
