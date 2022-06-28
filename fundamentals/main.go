package main

import (
	"fmt"
)

func newCard() string {
	return "Ace of spades"
}

func main() {
	// variables
	//var card string = "Ace of spades" // type declared
	//card := "Ace of spades"   // type infer initialization
	//card = "Five of diamonds" // replacing variable
	card := newCard()
	fmt.Println(card)

	// lists
	// array -> fixed
	// slice -> can grow

	cards := []string{"card1", "card2", "card3"}
	cards = append(cards, "card4")
	fmt.Println(cards)

	// iterate
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// change type
	greeting := "Hi there"
	fmt.Print([]byte(greeting))
}
