package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("We are learning JSON")

	// Original struct
	person := Person{Name: "John", Age: 34, IsAdult: true}

	// Marshal (encode) struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("person Data in JSON:", string(jsonData))

	// Unmarshal (decode) JSON back to struct
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("person Data after Unmarshalling:", personData)
}
