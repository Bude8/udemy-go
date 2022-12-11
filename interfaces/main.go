package main

import "fmt"

// Naive code demonstrating the uses of interfaces
// English and Spanish chat bot, both with a getGreeting() and printGreeting()
// function. The latter has identical logic, so should not need to be repeated

// Notes
// Interfaces are not 'generic' types
// Interfaces are 'implicit' - we don't manually have to say that our custom type satisfies some interface
// Interfaces are a contract to help us manage types - GARBAGE IN -> GARBAGE OUT. If our custom type's implementation of a function is broken, then interfaces won't help us!

// To whom it may concern
// We have a new type called 'bot'
// If you are a type in this program with a function called getGreeting
// AND you return a string
// You are an honourary member of the type 'bot', so you can now call the func
// printGreeting

// Everything in the interface must be satisfied to use the interface
type bot interface {
	getGreeting() string
	// getGreeting (string, int) (string, int)
	// funciton name, list of arg types, list of return types
}

type englishBot struct{}
type spanishBot struct{}

// Concrete types e.g. map, struct, englishBot - can create a value, change it, create new copies, etc.
// Interface types e.g. bot - we cannot create a value with type bot directly

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating English greeting
	return "Hello there!"
}

// We can omit the value sb if it's not used
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
