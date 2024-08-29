package main

import (
	"fmt"
)

var filename = "deck.json"

func main() {
	var deck Deck
	// deck = newDeck()
	// err := deck.saveToFile(filename)
	
	deck = readFromFile(filename)
	deck.shuffle()

	cardsOnHand, remaining := deal(deck, 4)
	fmt.Println("On hand:")
	printCards(cardsOnHand)
	fmt.Println("Remaining:")
	printCards(remaining)
}

func printCards(cards []string) {
	for i, card := range cards {
		fmt.Println(i + 1, card)
	}
}

