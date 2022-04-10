package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/mateoradman/go-cards/internal/models"
)

func createNewCompleteDeck() models.Deck {
	allCards := models.GetAllCards()
	deck := models.NewDeck(allCards)
	InMemoryStorage[deck.Deck_ID] = &deck
	return deck
}

func TestPOSTNewCompleteDeck(t *testing.T) {
	application := &application{}
	request, _ := http.NewRequest("POST", "/api/v1/decks", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{}
	application.createDeck(response, request, params)
	minimalDeck := models.MinimalDeck{}
	err := json.NewDecoder(response.Body).Decode(&minimalDeck)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if minimalDeck.Deck_ID == "" {
		t.Errorf("Deck ID is empty")
	}
	if minimalDeck.Shuffled != false {
		t.Errorf("Deck is shuffled")
	}
	if minimalDeck.Remaining != 52 {
		t.Errorf("Deck does not contain 52 cards")
	}
}

func TestPOSTNewShuffledDeck(t *testing.T) {
	application := &application{}
	request, _ := http.NewRequest("POST", "/api/v1/decks?shuffled=true", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{}
	application.createDeck(response, request, params)
	minimalDeck := models.MinimalDeck{}
	err := json.NewDecoder(response.Body).Decode(&minimalDeck)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if minimalDeck.Shuffled != true {
		t.Errorf("Deck is not shuffled")
	}
}

func TestPOSTNewDeckWithCustomCards(t *testing.T) {
	application := &application{}
	request, _ := http.NewRequest("POST", "/api/v1/decks?cards=4S,8S,QS,JD,KH,AD", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{}
	application.createDeck(response, request, params)
	minimalDeck := models.MinimalDeck{}
	err := json.NewDecoder(response.Body).Decode(&minimalDeck)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if minimalDeck.Remaining != 6 {
		t.Errorf("Deck does not contain 6 cards as expected.")
	}
}

func TestGETDeck(t *testing.T) {
	newDeck := createNewCompleteDeck()
	application := &application{}
	// GET the created deck
	request, _ := http.NewRequest("GET", "api/v1/decks/:uuid", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{
		httprouter.Param{
			Key:   "uuid",
			Value: newDeck.Deck_ID,
		}}
	application.getDeck(response, request, params)
	deck := models.Deck{}
	err := json.NewDecoder(response.Body).Decode(&deck)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Validate the deck
	if deck.Deck_ID != newDeck.Deck_ID {
		t.Errorf("Deck ID is not the same as expected.")
	}
	if deck.Shuffled != newDeck.Shuffled {
		t.Errorf("Deck is shuffled")
	}
	if deck.Remaining != newDeck.Remaining {
		t.Errorf("Deck does not contain the expected number of cards")
	}
	if len(deck.Cards) != 52 {
		t.Errorf("Deck does not contain 52 cards")
	}
}

func TestGETDeckNotFound(t *testing.T) {
	application := &application{}
	request, _ := http.NewRequest("GET", "/api/v1/decks/some_id", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{}
	application.getDeck(response, request, params)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %v, got %v", http.StatusNotFound, response.Code)
	}
}

func TestPOSTDrawCardFromDeck(t *testing.T) {
	newDeck := createNewCompleteDeck()
	application := &application{}
	// Draw a card from the deck
	request, _ := http.NewRequest("POST", "/api/v1/decks/:uuid/draw?count=12", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{
		httprouter.Param{
			Key:   "uuid",
			Value: newDeck.Deck_ID,
		}}
	application.drawCardFromDeck(response, request, params)
	deckCards := models.DeckCards{}
	err := json.NewDecoder(response.Body).Decode(&deckCards)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	if len(deckCards.Cards) != 12 {
		t.Errorf("Number of drawn cards is incorrect. Expected %v, got %v", 12, len(deckCards.Cards))
	}
}

func TestPOSTDrawCardFromDeckCorrectRemainingNumber(t *testing.T) {
	newDeck := createNewCompleteDeck()
	application := &application{}
	// Draw a card from the deck
	request, _ := http.NewRequest("POST", "/api/v1/decks/:uuid/draw?count=12", nil)
	response := httptest.NewRecorder()
	params := httprouter.Params{
		httprouter.Param{
			Key:   "uuid",
			Value: newDeck.Deck_ID,
		}}
	application.drawCardFromDeck(response, request, params)
	inMemoryDeck := InMemoryStorage[newDeck.Deck_ID]
	if inMemoryDeck.Remaining != 40 {
		t.Errorf("Remaining number is incorrect. Expected %v, got %v", 40, inMemoryDeck.Remaining)
	}
}
