package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	fmt.Println(resp.Body)
	// Body is of type Reader, which is an interface used for every external input
	// http, text, binary, csv, image, etc
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// custom implementation of writer interface (Write method)
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
