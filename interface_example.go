package main

import "fmt"

type bot interface {
	getGreating() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreating() string {
	// VERY custom logic for generating an english greating
	return "Hi There!"
}

func (sb spanishBot) getGreating() string {
	return "Hola!"
}

func printGreating(b bot) {
	fmt.Println(b.getGreating())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb) // Hi There!
	printGreating(sb) // Hola!
}
