package models

import "testing"

func TestCard(t *testing.T) {
	card := NewCard("2", "S")
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

func TestGetAllCardsSequentialOrder(t *testing.T) {
	cards := GetAllCards()
	if cards[0].Value != "2" {
		t.Errorf("Expected cards[0].Value to be 2, got %s", cards[0].Value)
	}
	if cards[0].Suit != "SPADES" {
		t.Errorf("Expected cards[0].Suit to be SPADES, got %s", cards[0].Suit)
	}
	if cards[0].Code != "2S" {
		t.Errorf("Expected cards[0].Code to be 2S, got %s", cards[0].Code)
	}
	if cards[1].Value != "3" {
		t.Errorf("Expected cards[1].Value to be 3, got %s", cards[1].Value)
	}
	if cards[1].Suit != "SPADES" {
		t.Errorf("Expected cards[1].Suit to be SPADES, got %s", cards[1].Suit)
	}
	if cards[1].Code != "3S" {
		t.Errorf("Expected cards[1].Code to be 3S, got %s", cards[1].Code)
	}
	if cards[2].Value != "4" {
		t.Errorf("Expected cards[2].Value to be 4, got %s", cards[2].Value)
	}
	if cards[2].Suit != "SPADES" {
		t.Errorf("Expected cards[2].Suit to be SPADES, got %s", cards[2].Suit)
	}
	if cards[2].Code != "4S" {
		t.Errorf("Expected cards[2].Code to be 4S, got %s", cards[2].Code)
	}
	if cards[3].Value != "5" {
		t.Errorf("Expected cards[3].Value to be 5, got %s", cards[3].Value)
	}
	if cards[51].Suit != "HEARTS" {
		t.Errorf("Expected cards[-1].Suit to be HEARTS, got %s", cards[51].Suit)
	}
	if cards[51].Code != "AH" {
		t.Errorf("Expected cards[-1].Code to be AH, got %s", cards[51].Code)
	}
	if cards[51].Value != "ACE" {
		t.Errorf("Expected cards[-1].Value to be ACE, got %s", cards[51].Value)
	}
}

func TestNewCardFromCode(t *testing.T) {
	card := NewCardFromCode("2S")
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

func TestNewCardsFromCodes(t *testing.T) {
	cards := NewCardsFromCodes([]string{"2S", "3S"})
	if len(cards) != 2 {
		t.Errorf("Expected len(cards) to be 2, got %d", len(cards))
	}
	if cards[0].Value != "2" {
		t.Errorf("Expected cards[0].Value to be 2, got %s", cards[0].Value)
	}
	if cards[0].Suit != "SPADES" {
		t.Errorf("Expected cards[0].Suit to be SPADES, got %s", cards[0].Suit)
	}
	if cards[0].Code != "2S" {
		t.Errorf("Expected cards[0].Code to be 2S, got %s", cards[0].Code)
	}
	if cards[1].Value != "3" {
		t.Errorf("Expected cards[1].Value to be 3, got %s", cards[1].Value)
	}
	if cards[1].Suit != "SPADES" {
		t.Errorf("Expected cards[1].Suit to be SPADES, got %s", cards[1].Suit)
	}
	if cards[1].Code != "3S" {
		t.Errorf("Expected cards[1].Code to be 3S, got %s", cards[1].Code)
	}
}
