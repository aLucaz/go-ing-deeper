package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardNumbers := []string{
		"one", "two", "three", "four",
		"five", "six", "seven", "eight",
		"nine", "ten", "eleven", "twelve",
		"thirteen",
	}
	cardIcons := []string{
		"hearts",
		"spades",
		"trevols",
		"diamonds",
	}
	for _, cardIcon := range cardIcons {
		for _, cardNumber := range cardNumbers {
			cards = append(cards, cardNumber+" of "+cardIcon)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(currentDeck deck, handSize int) (deck, deck) {
	return currentDeck[:handSize], currentDeck[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR occurred reading file")
		fmt.Println(err)
		os.Exit(-1)
	}
	return strings.Split(string(file), ",")
}
