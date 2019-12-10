package main

func main() {
	cards := deck{"Queens", getCard()}
	cards = append(cards, "Kings")
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
