package main

import "fmt"

func main() {
	deck := newDeck()
	cards, remaining := deal(deck, 4)
	fmt.Println("On hand:")
	printCards(cards)
	fmt.Println("Remaining:")
	printCards(remaining)
}

func printCards(cards []string) {
	for i, card := range cards {
		fmt.Println(i + 1, card)
	}
}

