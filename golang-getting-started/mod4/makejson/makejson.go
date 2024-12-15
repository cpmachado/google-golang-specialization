package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Prompt for name
	var name string
	fmt.Print("Name: ")
	if _, err := fmt.Scanln(&name); err != nil {
		log.Fatal(err)
	}

	// Prompt for address
	var address string
	fmt.Print("Address: ")
	if _, err := fmt.Scanln(&address); err != nil {
		log.Fatal(err)
	}

	// Create map, and insert data
	user := make(map[string]string)
	user["name"] = name
	user["address"] = address

	// Serialize map
	s, _ := json.Marshal(user)
	// Print map
	fmt.Println(string(s))
}
