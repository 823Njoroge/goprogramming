package main

import (
	"encoding/json"
	"fmt"
)

// EPL struct with JSON tags for encoding
type EPL struct {
	Name     string `json:"name"`
	Trophies int    `json:"trophies"`
	Players  int    `json:"players"`
}

func main() {
	fmt.Println("JSON encoding in Go")
	encodeJSON()
}

// encodeJSON function handles struct to JSON encoding
func encodeJSON() {
	league := []EPL{
		{"Chelsea", 23, 56},
		{"Manjesta", 12, 67},
		{"Liverpool", 9, 25},
		{"Arsenal", 12, 45},
	}

	// Print the struct array
	fmt.Println("Original Struct Data:", league)

	// Encode to JSON with indentation
	finalJSON, err := json.MarshalIndent(league, "", "\t")
	if err != nil {
		panic(err)
	}

	// Print the encoded JSON
	fmt.Printf("Encoded JSON:\n%s\n", finalJSON)
}
