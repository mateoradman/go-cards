package models

import "testing"

func TestCard(t *testing.T) {
	card := Card{}
	card.NewCard("2", "SPADES")
	if card.Value != "2" {
		t.Errorf("Expected card.Value to be 2, got %s", card.Value)
	}
	if card.Suit != "SPADES" {
		t.Errorf("Expected card.Suit to be SPADES, got %s", card.Suit)
	}
	if card.Code != "2S" {
		t.Errorf("Expected card.Code to be 2S, got %s", card.Code)
	}
}

func TestGetAllCards(t *testing.T) {
	cards := GetAllCards()
	if len(cards) != 52 {
		t.Errorf("Expected len(cards) to be 52, got %d", len(cards))
	}
}
