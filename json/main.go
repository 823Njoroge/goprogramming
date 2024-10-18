package main

import (
	"encoding/json"
	"fmt"
)

type EPL struct {
	Name     string `json:"name"`
	Trophies int    `json:"trophies"`
	Players  int    `json:"players"`
}

func main() {
	fmt.Println("JSON encoding in Go")
	encodeJSON()
}
//json data
func encodeJSON() {
	league := []EPL{
		{"Chelsea", 23, 56},
		{"Manjesta", 12, 67},
		{"Liverpool", 9, 25},
		{"Arsenal", 12, 45},
	}

	// Printing the struct array for reference
	fmt.Println("Original Struct Data:", league)

	// Encoding to 
	finalJSON, err := json.MarshalIndent(league, "", "\t")
	if err != nil {
		panic(err)
	}
	

	// Printing the encoded JSON
	fmt.Printf("Encoded JSON:\n%s\n", finalJSON)
}
