package models

var CardsData = map[string]string{
	"A":  "ACE",
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"J":  "JACK",
	"Q":  "QUEEN",
	"K":  "KING",
}

var Suits = map[string]string{
	"S": "SPADES",
	"H": "HEARTS",
	"D": "DIAMONDS",
	"C": "CLUBS",
}

type Card struct {
	Value string `json:"value"` // Value of the card
	Suit  string `json:"suit"`  // Suit of the card
	Code  string `json:"code"`  // Code of the card
}

func NewCard(value string, suit string) Card {
	card := Card{}
	card.Value = CardsData[value]
	card.Suit = Suits[suit]
	card.Code = value + suit
	return card
}

func GetAllCards() []Card {
	var cards []Card
	for value := range CardsData {
		for suit := range Suits {
			card := NewCard(value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

func NewCardFromCode(code string) Card {
	card := Card{}
	card.Code = code
	card.Value = CardsData[code[:1]]
	card.Suit = Suits[code[1:]]
	return card
}

func NewCardsFromCodes(codes []string) []Card {
	var cards []Card
	for _, code := range codes {
		cards = append(cards, NewCardFromCode(code))
	}
	return cards
}
