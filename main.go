package main

func main() {
	cards := newDeck()
	cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// for i := 0; i < 2; i++ {
	// 	fmt.Println(cards[i])
	// }

}

func getCard() string {
	return "Ace"
}
