package main

func main() {
	cards := newDeck()

	// cards.print()
	// cards.pizazze()

	// deal returns 2 different deck values
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.pizazze()

}

func newCard() string {
	return "Five of Diamonds"
}
