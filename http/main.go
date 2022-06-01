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
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (w logWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	fmt.Println(string(p))
	fmt.Printf("Just wrote %d bytes to log\n\n", n)
	return n, nil
}
