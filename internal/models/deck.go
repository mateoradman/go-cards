package models

import (
	"math/rand"

	"github.com/google/uuid"
)

type Deck struct {
	Deck_ID   string `json:"deck_id"`   // Unique identifier for the deck
	Shuffled  bool   `json:"shuffled"`  // Whether the deck has been shuffled
	Remaining int    `json:"remaining"` // Number of cards remaining in the deck
	Cards     []Card `json:"cards"`     // Cards in the deck
}

func (d *Deck) NewDeck(cards []Card) {
	d.Deck_ID = uuid.New().String() // Generate a new unique ID for the deck
	d.Shuffled = false
	d.Remaining = len(cards)
	d.Cards = make([]Card, d.Remaining)
	for i, card := range cards {
		d.Cards[i] = card
	}
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	d.Shuffled = true
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	d.Remaining = len(d.Cards)
	return card
}
