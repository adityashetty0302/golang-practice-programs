package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected 16, got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, got %v", cards[len(cards)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	cards := newDeck()
	cards.saveToFile("_decktesting.txt")
	newCards := newDeckFromFile("_decktesting.txt")

	if len(newCards) != 16 {
		t.Errorf("Expected 16, got %v", len(newCards))
	}
	os.Remove("_decktesting.txt")
}
