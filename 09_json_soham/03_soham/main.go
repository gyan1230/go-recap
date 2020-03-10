package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJSON := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"legs":2,"animals":"none"}`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(birdJSON), &result)
	if err != nil {
		fmt.Println(err)
	}

	birds := result["birds"].(map[string]interface{})
	for k, v := range birds {
		fmt.Println(k, ":", v.(string))
	}
	leg := result["legs"].(float64)
	fmt.Println("legs :", leg)
	animal := result["animals"].(string)
	fmt.Println("Is animal:", animal)
}
