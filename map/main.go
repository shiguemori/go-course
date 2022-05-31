package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
		"black": "#000000",
		"blue":  "#0000ff",
	}
	printMap(colors)
}

func printMap(colors map[string]string) {
	for key, value := range colors {
		fmt.Println("Hex code for", key, "is", value)
	}
}
