package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	err := cards.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error saving cards to file", err)
		return
	}

	//deckFromFile := newDeckFromFile("my_cards")
	//deckFromFile.print()
}
