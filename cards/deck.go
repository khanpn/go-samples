package main

import "fmt"

type Deck struct {
	cards []string
}

func (d Deck) print() {
	for i, card := range d.cards {
		fmt.Println(i, card)
	}
}

func newDeck() Deck {
	d := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d.cards = append(d.cards, value + " of " + suit)
		}
	}

	return d
}

