package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Name struct {
	fname string
	lname string
}

const (
	nameMaxRuneLength = 20
)

func main() {
	// Prompt and read filename
	var filename string

	fmt.Print("Enter a filename: ")
	if _, err := fmt.Scan(&filename); err != nil {
		log.Fatal(err)
	}

	// Open file
	filehandle, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

	// Create buffered scanner by line
	reader := bufio.NewScanner(filehandle)
	reader.Split(bufio.ScanLines)
	names := make([]Name, 0, 10)

	// Read line by line
	for reader.Scan() {
		fields := strings.Split(reader.Text(), " ")
		if len(fields) != 2 {
			log.Fatal(errors.New("Invalid number of fields"))
		}

		// Test for maximum size
		for _, field := range fields {
			if n := utf8.RuneCountInString(field); n > nameMaxRuneLength {
				log.Fatalf("'%s' has '%d' runes, bigger than limit", field, n)
			}
		}

		// Apend to slice
		names = append(names, Name{fields[0], fields[1]})
	}

	for _, name := range names {
		fmt.Printf("fname: '%s', lname: '%s'\n", name.fname, name.lname)
	}
}
