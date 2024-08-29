package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if (len(d.Cards) != 16) {
		t.Errorf("Expected deck length of 16, but got %d", len(d.Cards))
	}

	if (d.Cards[0] != "Ace of Spades") {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d.Cards[0])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck();
	deck.saveToFile(filename)

	loadedDect := readFromFile(filename)

	if (len(loadedDect.Cards) != 16) {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDect.Cards))
	}

	os.Remove(filename)
}