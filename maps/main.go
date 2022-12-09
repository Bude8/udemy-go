package main

import "fmt"

// Maps!
// Similar to structs, but the keys AND values are statically typed - so they are the same type

func main() {
	// var colours map[string]string
	// colours := make(map[string]string)

	colours := map[string]string{
		"red":   "#ff0000",
		"foo":   "#4bf123",
		"white": "#ffffff",
	}

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
