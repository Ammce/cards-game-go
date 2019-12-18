package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Error happened, expecting 20 cards, but got %v", len(d))
	}
	if d[0] != "Ace of Diamond" {
		t.Errorf("Error Happened, expecting Ace of Diamond, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Error happened, expecting Four of Club, bug tog %v", d[len(d)-1])
	}
}

func TestSaveToFileAndGetCardsFromFile(t *testing.T) {
	fileName := "_deckTesting"
	os.Remove(fileName)
	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := getCardsFromFile(fileName)

	if len(loadedDeck) != 20 {
		t.Errorf("Error Happened, Expected 20 cards, but got %v", len(loadedDeck))
	}
	os.Remove(fileName)
}
