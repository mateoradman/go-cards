package models

var CardsData = map[string]string{
	"ACE":   "ACE",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"10":    "10",
	"JACK":  "JACK",
	"QUEEN": "QUEEN",
	"KING":  "KING",
}

var Suits = map[string]string{
	"SPADES":   "S",
	"HEARTS":   "H",
	"DIAMONDS": "D",
	"CLUBS":    "C",
}

type Card struct {
	Value string `json:"value"` // Value of the card
	Suit  string `json:"suit"`  // Suit of the card
	Code  string `json:"code"`  // Code of the card
}

func (c *Card) NewCard(value string, suit string) {
	c.Value = value
	c.Suit = suit
	c.Code = CardsData[value] + Suits[suit]
}

func GetAllCards() []Card {
	var cards []Card
	for value := range CardsData {
		for suit := range Suits {
			card := Card{}
			card.NewCard(value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}
