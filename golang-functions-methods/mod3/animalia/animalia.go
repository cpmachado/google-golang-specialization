package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Animal represents an animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints what a given animal eats
func (a *Animal) Eat() {
	fmt.Printf("It eats %s\n", a.food)
}

// Move prints how a given animal moves
func (a *Animal) Move() {
	fmt.Printf("It %ss\n", a.locomotion)
}

// Move prints how a given animal speaks
func (a *Animal) Speak() {
	fmt.Printf("It does '%s'\n", a.noise)
}

var animals = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func help() {
	fmt.Println("Commands are in the format: <animal name> <action>")
	fmt.Println("\taction being one of ('eat', 'move', 'speak')")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read input
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Break up tokens
		tokens := strings.Split(strings.TrimSpace(s), " ")
		if len(tokens) != 2 {
			help()
			continue
		}
		name := strings.ToLower(tokens[0])
		cmd := strings.ToLower(tokens[1])

		// Check if it exists
		animal, found := animals[name]
		if !found {
			fmt.Printf("Animal '%s' not found\n", tokens[0])
			continue
		}

		// Process Request
		switch cmd {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			help()
		}
	}
}
