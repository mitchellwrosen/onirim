package main

import (
	"testing"
)

func TestNewBasicDeck(t *testing.T) {
	deck := NewBasicDeck()

	if len(deck) != 76 {
		t.Error("New basic deck contains 76 cards")
	}
}
