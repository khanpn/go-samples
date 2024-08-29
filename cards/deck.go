package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Deck struct {
	Cards []string `json:"cards"`
}

func (d Deck) print() {
	for i, card := range d.Cards {
		fmt.Println(i + 1, card)
	}
}

func newDeck() Deck {
	d := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d.Cards = append(d.Cards, value + " of " + suit)
		}
	}

	return d
}

func deal(d Deck, handSize int) ([]string, []string) {
	return d.Cards[:handSize], d.Cards[handSize:]
}

func (d Deck) toByteSlices() []byte {
	bytes, _ := json.Marshal(d);
	return bytes
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlices(), os.ModePerm)
}

func readFromFile(filename string) Deck {
	var deck Deck
	bytes, err := os.ReadFile(filename)
	
	if (err == nil) {
		err = json.Unmarshal(bytes, &deck)
	}
	
	if (err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck
}

func (d Deck) shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}