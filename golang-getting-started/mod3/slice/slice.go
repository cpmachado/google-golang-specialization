package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create empty integer slice
	s := make([]int, 0, 3)
	input := ""

	// Readline
	fmt.Print("Enter an integer('x' to quit): ")
	_, err := fmt.Scanln(&input)

	// conditions to continue, buffer isn't empty or X
	for err == nil && strings.ToLower(input) != "x" {

		// convert integer
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		// Append to slice
		s = append(s, n)

		// Sort
		sort.Ints(s)
		fmt.Println(s)
		fmt.Print("Enter an integer('x' to quit): ")
		_, err = fmt.Scanln(&input)
	}
}
