package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	// in structs we cannot iterate oveer the key value pais
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
	}
	colors["black"] = "#000000"
	fmt.Println(colors)
	delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)
}
