package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	// if d[0] != "Ace of Spades" {
	// 	t.Errorf("Expected Ace of Spades and got '%v' ", d[0])
	// }

}
