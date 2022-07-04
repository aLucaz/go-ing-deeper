package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://tiendamia.com",
	}
	c := make(chan string)
	for _, link := range links {
		// to let the scheduler decide when to add go routines
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down!")
		c <- "Link is down"
		return
	}
	fmt.Println(link, " is up!")
	c <- "Link is up"
}
