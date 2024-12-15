package main

import (
	"fmt"
	"io"
	"log"
)

// Swap Swaps the values of a given index i and i+1 in a slice of ints
func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

// BubbleSort sorts a given slice s of ints
func BubbleSort(s []int) {
	n := len(s)
	swapped := true
	// Should iterate from 1
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				Swap(s, i)
				swapped = true
			}
		}
		n--
	}
}

func main() {
	var i int
	s := make([]int, 10)

	fmt.Print("Enter 10 integers: ")

	// conditions to continue, buffer isn't empty or X
	for i = 0; i < 10; i++ {
		// Using buffered io, thus is can't break in the middle of processing a
		// line with numbers. E.g.: 1 2\n3 EOF
		if _, err := fmt.Scanf("%d", &s[i]); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	s = s[:i]
	BubbleSort(s)
	fmt.Printf("%v\n", s)
}
