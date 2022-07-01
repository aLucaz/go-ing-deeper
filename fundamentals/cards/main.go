package main

import "fmt"

// go run main.go deck.go

func main() {
	cards := newDeck()
	cards.print()
	hand, rest := deal(cards, 5)
	hand.print()
	rest.print()
	handString := hand.toString()
	fmt.Println(handString)
	err := hand.saveToFile("myhand.txt")
	if err != nil {
		return
	}
	myHand := newDeckFromFile("myhand.txt")
	myHand.print()
	myHand.shuffle()
	myHand.print()
}
