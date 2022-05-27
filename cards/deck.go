package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const folder = "my-decks"

type deck []string

// function to create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function to shuffle the deck
func (d *deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range *d {
		newPosition := r.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPosition] = (*d)[newPosition], (*d)[i]
	}
}

// function to pick cards from the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to print the deck
func (d *deck) print() {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<", "Printing deck : ", ">>>>>>>>>>>>>>>>>>>>>>>>>")
	for _, card := range *d {
		fmt.Println(card)
	}
}

// function to verify the folder
func verifyFolder() error {
	newpath := filepath.Join(".", folder)
	return os.MkdirAll(newpath, os.ModePerm)
}

// function to set the file name
func setFileName(filename string) string {
	return folder + "/" + filename + ".txt"
}

// function to save the deck to a file
func (d *deck) saveToFile(filename string) error {
	err := verifyFolder()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	p := filepath.FromSlash(setFileName(filename))
	file, err := os.Create(p)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			println("Error closing file")
		}
	}(file)

	for _, card := range *d {
		_, err = file.WriteString(card + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// function to load a deck from a file
func newDeckFromFile(filename string) deck {
	p := filepath.FromSlash(setFileName(filename))
	file, err := os.Open(p)
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
			os.Exit(1)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	return lines
}
