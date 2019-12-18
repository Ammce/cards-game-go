package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Error happened, expecting 20 cards, but got %v", len(d))
	}
}
