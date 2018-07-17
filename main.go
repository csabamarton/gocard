package main

import
	"fmt"


var createDeck = newDeckFromFile

func main() {
	cards := createDeck("my_cards")

	remainingCards, _ := deal(cards, 5)

	 remainingCards.print()

	remainingCards.shuffle()

	fmt.Println(remainingCards.toString())

	cards.saveToFile("my_cards2")
}
