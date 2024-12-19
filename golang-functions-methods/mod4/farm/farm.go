package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Animal  is an interface for animals
type Animal interface {
	// Eat prints what a given animal eats
	Eat()
	// Move prints how a given animal moves
	Move()
	// Move prints how a given animal speaks
	Speak()
}

type Cow struct{}

// Eat prints what a given cow eats
func (_ *Cow) Eat() {
	fmt.Println("grass")
}

// Move prints how a given cow moves
func (_ *Cow) Move() {
	fmt.Println("walk")
}

// Move prints how a given cow speaks
func (_ *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

// Eat prints what a given Bird eats
func (_ *Bird) Eat() {
	fmt.Println("worms")
}

// Move prints how a given Bird moves
func (_ *Bird) Move() {
	fmt.Println("fly")
}

// Move prints how a given Bird speaks
func (_ *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

// Eat prints what a given Snake eats
func (_ *Snake) Eat() {
	fmt.Println("mice")
}

// Move prints how a given Snake moves
func (_ *Snake) Move() {
	fmt.Println("slither")
}

// Move prints how a given Snake speaks
func (_ *Snake) Speak() {
	fmt.Println("hsss")
}

func help() {
	fmt.Println("Commands are in the format:")
	fmt.Println(" - query <animal name> <action>")
	fmt.Println(" - newanimal <animal name> <species>")
	fmt.Println("\taction being one of ('eat', 'move', 'speak')")
	fmt.Println("\tspecies being one of ('cow', 'bird', 'snake')")
}

// AddAnimal Attempts to add animal and returns success
func AddAnimal(animals map[string]Animal, name, species string) bool {
	switch species {
	case "cow":
		animals[name] = &Cow{}
	case "bird":
		animals[name] = &Bird{}
	case "snake":
		animals[name] = &Snake{}
	default:
		return false
	}
	fmt.Println("Created it!")
	return true
}

// ProcessAction acts according to designated action and returns success
func ProcessAction(a Animal, action string) bool {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		return false
	}
	return true
}

func main() {
	animals := make(map[string]Animal)
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
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		tokens := strings.Split(s, " ")
		if len(tokens) != 3 {
			help()
			continue
		}

		cmd := strings.ToLower(tokens[0])
		name := strings.ToLower(tokens[1])
		switch cmd {
		case "query":
			action := strings.ToLower(tokens[2])
			// Check if it exists
			animal, found := animals[name]
			if !found {
				fmt.Printf("Animal '%s' not found\n", name)
				break
			}
			if !ProcessAction(animal, action) {
				help()
			}
		case "newanimal":
			species := strings.ToLower(tokens[2])
			if !AddAnimal(animals, name, species) {
				help()
			}
		default:
			help()
		}

	}
}
