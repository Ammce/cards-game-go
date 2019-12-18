package main

func main() {
	// cards := newDeck()
	// cards.print()
	// cars := []string{"Audi", "BMW", "Mercedes", "Opel", "Fiat"}
	// fmt.Println(cars[3:])
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// for i := 0; i < 2; i++ {
	// 	fmt.Println(cards[i])
	// }

	// hand, remainingCars := deal(cards, 5)
	// hand.print()
	// remainingCars.print()

	// hearts := retriveHearts(cards)
	// hearts.print()

	// cards.saveToFile("myFile.txt")

	// numbers := []int{0, 1, 2, 3, 4, 5}
	// rand.Shuffle(len(numbers), func(i, j int) {
	// 	numbers[i], numbers[j] = numbers[j], numbers[i]
	// })
	// fmt.Println(numbers)
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func getCard() string {
	return "Ace"
}
