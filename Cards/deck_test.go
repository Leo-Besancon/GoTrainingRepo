package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 elements, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First element expected to be Ace of Spades, but was %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Last element expected to be King of Clubs, but was %v", len(d)-1)
	}

}

func TestSaveAndLoad(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveDeckToFile("_decktesting")

	d2 := newDeckFromFile("_decktesting")

	if len(d) != len(d2) {
		t.Errorf("The two decks don't contain the same number of cards!")
	}

}
