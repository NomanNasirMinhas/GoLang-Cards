package main

func main() {
	// cards := newDeck()
	// hand, remaining := deal(cards, 5)
	// fmt.Println("Hand Deck")
	// hand.print()
	// fmt.Println("Remaining")
	// remaining.print()
	// fmt.Println(cards.toString())
	cards := deckFromFile("my cards.txt")
	cards.shuffleDeck()
	cards.print()
}
