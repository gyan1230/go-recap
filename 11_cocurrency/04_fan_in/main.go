package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go producer(in, 100*time.Millisecond)
	go producer(in, 200*time.Millisecond)

	go reader(out)

	for i := range in {
		out <- i
	}
}

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}
