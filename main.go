package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println(remainingCards.toString())

	cards.saveToFile("my_cards2")
}


