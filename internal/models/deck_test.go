package models

import "testing"

var sampleCards = []Card{
	{Value: "2", Suit: "SPADES", Code: "2S"},
	{Value: "6", Suit: "HEARTS", Code: "6H"},
}

func TestNewCompleteDeck(t *testing.T) {
	cards := GetAllCards()
	deck := NewDeck(cards)
	if len(deck.Cards) != 52 {
		t.Errorf("Expected len(deck.Cards) to be 52, got %d", len(deck.Cards))
	}
}

func TestNewCustomDeck(t *testing.T) {
	deck := NewDeck(sampleCards)
	if len(deck.Cards) != 2 {
		t.Errorf("Expected len(deck.Cards) to be 2, got %d", len(deck.Cards))
	}
}

func TestShuffle(t *testing.T) {
	cards := GetAllCards()
	deck := NewDeck(cards)
	deck.Shuffle()
	if deck.Shuffled != true {
		t.Errorf("Expected deck.Shuffled to be true, got %t", deck.Shuffled)
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck(sampleCards)
	deck.Shuffle()
	cardsDrawn := deck.Draw(1)
	if len(cardsDrawn) != 1 {
		t.Errorf("Expected len(cards) to be 1, got %d", len(cardsDrawn))
	}
	if deck.Remaining != 1 {
		t.Errorf("Expected deck.Remaining to be 1, got %d", deck.Remaining)
	}
}

func TestDrawMoreThanRemaining(t *testing.T) {
	deck := NewDeck(sampleCards)
	deck.Shuffle()
	cardsDrawn := deck.Draw(3)
	if len(cardsDrawn) != 2 {
		t.Errorf("Expected len(cards) to be 2, got %d", len(cardsDrawn))
	}
	if deck.Remaining != 0 {
		t.Errorf("Expected deck.Remaining to be 0, got %d", deck.Remaining)
	}
}

func TestFilterDuplicateCards(t *testing.T) {
	cards := []Card{
		{Value: "2", Suit: "SPADES"},
		{Value: "6", Suit: "HEARTS"},
		{Value: "2", Suit: "SPADES"},
	}
	deck := NewDeck(cards)
	if len(deck.Cards) != 2 {
		t.Errorf("Expected len(deck.Cards) to be 2, got %d", len(deck.Cards))
	}
}

func TestCardExistsInDeck(t *testing.T) {
	deck := NewDeck(sampleCards)
	cardExists := deck.CardExists(Card{Value: "2", Suit: "SPADES", Code: "2S"})
	if cardExists != true {
		t.Errorf("Expected CardExistsInDeck(deck.Cards, Card{Value: \"2\", Suit: \"SPADES\"}) to be true, got %t", cardExists)
	}
}

func TestDeckToJSON(t *testing.T) {
	deck := NewDeck(sampleCards)
	json := deck.ToJSON()
	expected := `{"deck_id":"` + deck.Deck_ID + `","shuffled":false,"remaining":2,"cards":[{"value":"2","suit":"SPADES","code":"2S"},{"value":"6","suit":"HEARTS","code":"6H"}]}`
	if string(json) != expected {
		t.Errorf("Expected string(json) to be %s, got %s", expected, string(json))
	}
}

func TestMinimalDeck(t *testing.T) {
	deck := NewDeck(sampleCards)
	minimalDeck := MinimalDeck{
		Deck_ID:   deck.Deck_ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
	json := minimalDeck.ToJSON()
	expected := `{"deck_id":"` + deck.Deck_ID + `","shuffled":false,"remaining":2}`

	if string(json) != expected {
		t.Errorf("Expected string(json) to be %s, got %s", expected, string(json))
	}
}

func TestDeckCards(t *testing.T) {
	deck := NewDeck(sampleCards)
	deckCards := DeckCards{
		Cards: deck.Cards,
	}
	json := deckCards.ToJSON()
	expected := `{"cards":[{"value":"2","suit":"SPADES","code":"2S"},{"value":"6","suit":"HEARTS","code":"6H"}]}`
	if string(json) != expected {
		t.Errorf("Expected string(json) to be %s, got %s", expected, string(json))
	}
}
