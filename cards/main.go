package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("---")

	// hand, remainingCards := deal(cards, 4)
	// hand.print()
	// fmt.Println("---")
	// remainingCards.print()

	// cards.saveToFile("decks.txt")
	// cards = newDeckFromFile("decks.txt")
	// cards.print()

	// cards.shuffle()
	// cards.print()
}
