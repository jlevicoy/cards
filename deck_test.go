package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}

	if d[0] != "Ace of Espada" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])

	}

	if d[len(d)-1] != "Four of Treboles" {
		t.Errorf("Expected last card of Four of Club, but got %v", d[len(d)-1])

	}
}

func TestNewDeckfromFileSaveToFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckfromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

}
