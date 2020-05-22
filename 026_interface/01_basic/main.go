package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Length() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Length(); i++ {
		j := rand.Intn(s.Length() - i)
		s.Swap(i, j)
	}
}

type stringSlice []string

func (is stringSlice) Length() int {
	return len(is)
}

func (is stringSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	ss1 := stringSlice{"hello", "world", "Hi", "How", "Are", "You"}
	shuffle(ss1)
	fmt.Printf("%v\n", ss1)
}
