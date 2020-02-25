package main

import (
	"encoding/json"
	"fmt"
)

type bike struct {
	Model string `json:"model"`
	Color string `json:"color"`
	Cc    int    `json:"cc"`
}

type power struct {
	Displacement int `json:"CC"`
	Power        int `json:"PS"`
}

type car struct {
	Model string `json:"model"`
	Color string `json:"color"`
	Power power  `json:"power"`
}

func main() {
	//json raw data
	jsondata1 := `{"model":"ktm","color":"white","cc":200}`
	var b1 bike
	//unmarshal will convert json data to go struct pointed by struct variable
	err := json.Unmarshal([]byte(jsondata1), &b1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b1)
	//Marshal convert go struct to json data([]byte)
	b, err := json.MarshalIndent(b1, "\n", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	//same for array of json
	jsondata2 := `[{"model":"FZ","color":"blue","cc":150}, {"model":"BMW","color":"black","cc":310}]`

	var bikes []bike
	err = json.Unmarshal([]byte(jsondata2), &bikes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bikes)

	dis, err := json.MarshalIndent(bikes, "\n", "   ")
	fmt.Println(string(dis))

	//embedded objects
	cardata := `{"model":"polo","color":"beize","power":{"CC":1000,"PS":85}}`
	var c1 car
	json.Unmarshal([]byte(cardata), &c1)
	fmt.Println(c1)
	c, err := json.MarshalIndent(c1, "\n", "     ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(c))
}
