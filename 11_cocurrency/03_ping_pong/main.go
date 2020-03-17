package main

import (
	"fmt"
	"strconv"
	"time"
)

type ball struct {
	hits int
}

func main() {
	table := make(chan *ball)
	// go player("ping", table)
	// go player("pong", table)
	// go player("ding", table)
	// go player("dong", table)
	for i := 0; i < 99; i++ {
		go player(strconv.Itoa(i), table)
	}

	table <- &ball{hits: 1}
	time.Sleep(1 * time.Second)
	<-table
}

func player(player string, table chan *ball) {
	for {
		b := <-table
		b.hits++
		fmt.Println(player, b.hits)
		time.Sleep(100 * time.Millisecond)
		table <- b
	}
}
