package main

import (
	"testing"
)

func TestMainFunc(t *testing.T) {
	oldDeckCreation := createDeck
	// as we are exiting, revert newDeckFromFile back to oldDeckCreation at end of function
	defer func () { createDeck = oldDeckCreation }()
	createDeck = func (s string) (deck) {
		return deck{}
	}

	main()
}