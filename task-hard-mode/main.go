package main

import (
	"io"
	"os"
)

func main() {
	myFile := os.Args[0]
	file, err := os.OpenFile(myFile, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
