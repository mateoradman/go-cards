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
	// Keep track of the order of cards upon creation.
	// by default the deck is sequential: spades followed by diamonds, clubs, then hearts.
	suitsSortedKeys := []string{"S", "D", "C", "H"}
	cardsDataSortedKeys := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "K", "Q", "A"}

	var cards []Card
	for _, suit := range suitsSortedKeys {
		for _, value := range cardsDataSortedKeys {
			cards = append(cards, NewCard(value, suit))
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
