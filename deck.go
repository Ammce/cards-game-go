package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// func retriveHearts(d deck) deck {
// 	hearts := deck{}
// 	for _, singleCard := range d {
// 		splitedString := strings.Split(singleCard, " ")
// 		if splitedString[2] == "Diamond" {
// 			hearts = append(hearts, singleCard)
// 		}
// 	}
// 	return hearts
// }
