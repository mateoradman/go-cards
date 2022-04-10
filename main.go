package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mateoradman/go-cards/internal/models"
)

type application struct {
	logger *log.Logger
}

var InMemoryStorage = make(map[string]*models.Deck)

func main() {
	// Initialize a new logger which writes messages to the standard out stream,
	// prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{logger: logger}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	logger.Printf("Starting server on localhost%s", srv.Addr)
	logger.Fatal(srv.ListenAndServe())
}

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/decks/:uuid", app.getDeck)
	router.POST("/api/v1/decks/:uuid/draw", app.drawCardFromDeck)
	router.POST("/api/v1/decks", app.createDeck)
	return router
}

func (app *application) createDeck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	shuffledQuery := r.URL.Query().Get("shuffled")
	cardsQuery := r.URL.Query().Get("cards")
	var cards []models.Card
	if cardsQuery == "" {
		cards = models.GetAllCards()
	} else {
		// split cards string and create new card for each
		cardCodes := strings.Split(cardsQuery, ",")
		cards = models.NewCardsFromCodes(cardCodes)
	}
	deck := models.NewDeck(cards)

	if shuffledQuery == "true" {
		deck.Shuffle()
	}
	// Add deck to in memory storage
	InMemoryStorage[deck.Deck_ID] = &deck
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Convert the deck to JSON
	minimalDeck := models.MinimalDeck{Deck_ID: deck.Deck_ID, Shuffled: deck.Shuffled, Remaining: deck.Remaining}
	w.Write(minimalDeck.ToJSON())
}

func (app *application) getDeck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uuid := ps.ByName("uuid")
	deck, ok := InMemoryStorage[uuid]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deck.ToJSON())
}

func (app *application) drawCardFromDeck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uuid := ps.ByName("uuid")
	deck, ok := InMemoryStorage[uuid]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Validate count query param.
	count := r.URL.Query().Get("count")
	if count == "" {
		w.Write([]byte("Please specify the number of cards to draw."))
		return
	}
	countInt, err := strconv.Atoi(count)
	if err != nil {
		w.Write([]byte("Please specify the number of cards to draw as an integer."))
		return
	}
	if countInt > deck.Remaining {
		w.Write([]byte("Not enough cards in the deck."))
		return
	}

	drawnCards := deck.Draw(countInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert the cards to JSON
	deckCards := models.DeckCards{Cards: drawnCards}
	w.Write(deckCards.ToJSON())
}
