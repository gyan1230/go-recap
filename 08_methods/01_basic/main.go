package main

import (
	"fmt"
	"strings"
)

type car struct {
	Name, Model, Color string
	power              int
}

func main() {
	polo := car{"Polo", "Highline", "Flash red", 115}
	polo.details()
	fmt.Println(strings.Repeat("*", 30))
	polo.notmodified(85)
	polo.details()
	fmt.Println(strings.Repeat("*", 30))
	polo.modified(85)
	polo.details()

}

func (c car) details() {
	fmt.Println("Name of car  is: ", c.Name)
	fmt.Println("Model of car is: ", c.Model)
	fmt.Println("Color of car is: ", c.Color)
	fmt.Println("Power of car is: ", c.power)
}

func (c car) notmodified(pow int) {
	c.power = pow
}

func (c *car) modified(pow int) {
	c.power = pow
}
