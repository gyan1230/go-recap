package main

import "fmt"

type car struct {
	Name, Model, Color string
	Weight             float32
}

func main() {
	polo := car{
		Name:   "polo",
		Model:  "Highline plus",
		Color:  "beige",
		Weight: 1890.7,
	}

	fmt.Println("Model of car is", polo.Model)

	polopointer := &polo
	fmt.Println("Color of car is", polopointer.Color)

	fmt.Println(polopointer)

	vento := new(car)
	vento.Color = "blue"
	vento.Model = "comfortline"
	vento.Name = "vento"
	vento.Weight = 2430.32
	fmt.Println(vento)
}
