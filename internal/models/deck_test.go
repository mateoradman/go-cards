package models

import "testing"

func TestNewCompleteDeck(t *testing.T) {
	cards := GetAllCards()
	deck := Deck{}
	deck.NewDeck(cards)
	if len(deck.Cards) != 52 {
		t.Errorf("Expected len(deck.Cards) to be 52, got %d", len(deck.Cards))
	}
}

func TestNewCustomDeck(t *testing.T) {
	cards := []Card{
		{Value: "2", Suit: "SPADES"},
		{Value: "6", Suit: "HEARTS"},
	}
	deck := Deck{}
	deck.NewDeck(cards)
	if len(deck.Cards) != 2 {
		t.Errorf("Expected len(deck.Cards) to be 2, got %d", len(deck.Cards))
	}
}

func TestShuffle(t *testing.T) {
	cards := GetAllCards()
	deck := Deck{}
	deck.NewDeck(cards)
	deck.Shuffle()
	if deck.Shuffled != true {
		t.Errorf("Expected deck.Shuffled to be true, got %t", deck.Shuffled)
	}
	if deck.Cards[0].Value == "2" {
		t.Errorf("Expected deck.Cards[0].Value to be not 2, got %s", deck.Cards[0].Value)
	}
}
