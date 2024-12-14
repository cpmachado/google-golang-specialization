package main

import (
	"fmt"
	"log"
)

func main() {
	var a float64
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Truncated, the number is: %d\n", int64(a))
}
