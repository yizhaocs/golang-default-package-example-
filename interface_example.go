package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}

func printGreating(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreating(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greating
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
