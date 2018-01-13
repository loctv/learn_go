package main

import (
	"fmt"
)

type englishBot struct{}
type spainishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spainishBot{}
	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spainishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spainishBot) getGreeting() string {
	return "Hola!"
}
