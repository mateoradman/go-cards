package models

import (
	"encoding/json"
	"math/rand"

	"github.com/google/uuid"
)

func NewDeck(cards []Card) Deck {
	var d Deck
	d.Deck_ID = uuid.New().String() // Generate a new unique ID for the deck
	d.Shuffled = false
	d.Remaining = len(cards)
	// filter out duplicate cards
	d.Cards = RemoveDuplicateCards(cards)
	return d
}

func RemoveDuplicateCards(cards []Card) []Card {
	var filteredCards []Card
	for _, c := range cards {
		if !CardExists(c, filteredCards) {
			filteredCards = append(filteredCards, c)
		}
	}
	return filteredCards
}

func CardExists(card Card, cards []Card) bool {
	for _, c := range cards {
		if c.Value == card.Value && c.Suit == card.Suit {
			return true
		}
	}
	return false
}


type Deck struct {
	Deck_ID   string `json:"deck_id"`   // Unique identifier for the deck
	Shuffled  bool   `json:"shuffled"`  // Whether the deck has been shuffled
	Remaining int    `json:"remaining"` // Number of cards remaining in the deck
	Cards     []Card `json:"cards"`     // Cards in the deck
}

func (d *Deck) ToJSON() []byte {
	json, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return json
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
	d.Shuffled = true
}

func (d *Deck) Draw(count int) []Card {
	var drawnCards []Card
	for i := 0; i < count; i++ {
		if d.Remaining == 0 {
			break
		}
		drawnCards = append(drawnCards, d.Cards[0])
		d.Cards = d.Cards[1:]
		d.Remaining--
	}
	return drawnCards
}

func (d *Deck) CardExists(card Card) bool {
	for _, c := range d.Cards {
		if c.Value == card.Value && c.Suit == card.Suit {
			return true
		}
	}
	return false
}

type MinimalDeck struct {
	Deck_ID   string `json:"deck_id"`   // Unique identifier for the deck
	Shuffled  bool   `json:"shuffled"`  // Whether the deck has been shuffled
	Remaining int    `json:"remaining"` // Number of cards remaining in the deck
}

func (d *MinimalDeck) ToJSON() []byte {
	json, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return json
}

type DeckCards struct {
	Cards []Card `json:"cards"`
}

func (d *DeckCards) ToJSON() []byte {
	json, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return json
}
