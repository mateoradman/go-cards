# go-cards

Hello, user!
Welcome to the Cards API in Go. 


## How to run the API server?
From the root of the project run
```
make run
```

By default, the server runs on `localhost:8080`.
The server uses in-memory storage and all data is lost once the program terminates.

## API Endpoints

### 1. Create a new deck
This endpoint accepts a POST request which creates a new deck.
By default, the deck will be complete, amounting to 52 playing cards and it will be sequential (spades, followed by diamonds, clubs and hearts).
Example:
```
curl -X POST "localhost:8080/api/v1/decks"
```
Example response:
```json
{"deck_id":"b86f9666-fc6f-47ad-9c5b-71b23654d40e","shuffled":false,"remaining":52}
```
It is possible to create a shuffled deck by adding a request parameter:
```
curl -X POST "localhost:8080/api/v1/decks?shuffled=true"
```
Lastly, it is possible to create a custom deck by specifying the cards in the deck.
```
curl -X POST "localhost:8080/api/v1/decks?shuffled=true?cards=AS,KD,AC"
```


### 2. Open a deck
GET request to open a deck by UUID.
Example request:
```
curl "localhost:8080/api/v1/decks/b86f9666-fc6f-47ad-9c5b-71b23654d40e"
```

Example response:
```json
{
"deck_id":"b86f9666-fc6f-47ad-9c5b-71b23654d40e",
"shuffled":false,
"remaining":52,
"cards":[
    {"value":"2","suit":"SPADES","code":"2S"},
    {"value":"3","suit":"SPADES","code":"3S"},
    {"value":"4","suit":"SPADES","code":"4S"},
    ...]
}
```

### 3. Draw a card from a deck
Accepts a POST request to draw an X number of cards from the deck. Once the cards are drawn, they are not existing in the deck anymore.

Example request:
```
curl "localhost:8080/api/v1/decks/b86f9666-fc6f-47ad-9c5b-71b23654d40e/draw?count=2"
```

Example response:
```json
{
"cards":[
    {"value":"2","suit":"SPADES","code":"2S"},
    {"value":"3","suit":"SPADES","code":"3S"}]
}
```


## How to run tests?
From the root of the project run
```
make test
```

