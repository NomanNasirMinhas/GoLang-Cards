package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length 16, but got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected length Ace of Spades, but got %v", d[0])
	}
}
