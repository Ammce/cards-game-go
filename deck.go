package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func getCardsFromFile(fileName string) deck {
	response, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	}
	s := strings.Split(string(response), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
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
