package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* Findian validates ian criteria */
func Findian(s string) bool {
	s = strings.ToLower(s)
	return (strings.HasPrefix(s, "i") &&
		strings.HasSuffix(s, "n") &&
		strings.ContainsRune(s, 'a'))
}

func main() {
	fmt.Printf("Enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	/* Remove newline */
	s = strings.TrimSpace(s)

	if Findian(s) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
