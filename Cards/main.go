package main

func main() {
	cards := newDeck()

	cards.print()
	// cards.pizazze()
}

func newCard() string {
	return "Five of Diamonds"
}
