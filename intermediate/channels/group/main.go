package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://tiendamia.com",
	}
	var wg sync.WaitGroup
	for _, web := range links {
		lRefCopy := web
		go func(link string) {
			defer wg.Done() // the signal that this function is over
			time.Sleep(time.Second * 1)
			checkLink(link)
		}(lRefCopy)
		wg.Add(1) // adding a new routine
	}
	wg.Wait() // wait all the routines to give Done
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down!")
		return
	}
	fmt.Println(link, " is up!")
}
