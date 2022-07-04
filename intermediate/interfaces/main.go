package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	name string
}

type spanishBot struct {
	name string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola ameos"
}

func main() {
	bobBot := englishBot{
		name: "bob",
	}

	carlosBot := spanishBot{
		name: "carlos",
	}

	printGreeting(bobBot)
	printGreeting(carlosBot)
}
