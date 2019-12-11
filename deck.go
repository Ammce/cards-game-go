package main

import "fmt"

type deck []string

func newDeck() deck {
	cardDeck := deck{}

	cardSuits := []string{"Diamond", "Heart", "Spades", "Club"}

	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, card := range cardValues {
			cardDeck = append(cardDeck, card+" of "+suit)
		}
	}
	return cardDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
