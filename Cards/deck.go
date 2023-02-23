package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cS := range cardSuits {
		for _, cV := range cardValues {
			cards = append(cards, fmt.Sprintln(cV, " of ", cS))
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) pizazze() {
	for i, card := range d {
		fmt.Println("Index: ", i, " The card is a(n) ", card)
	}
}
