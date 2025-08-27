package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Card struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	SetCode  string `json:"setCode"`
	Quantity int    `json:"quantity"`
}

var cards = map[string]Card{}

func serverStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server up and runnning...\n"))
}

// STEP 1
// POST (from App.vue form) to create GET Scryfall request (step 2)
func createCard(w http.ResponseWriter, r *http.Request) {
	var card Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	if card.ID == "" {
		card.ID = strconv.Itoa(len(cards) + 1)
	}

	// DELETE?
	// _, exists := cards[card.ID]
	// if exists {
	// 	http.Error(w, "Card already exists", http.StatusConflict)
	// 	return
	// }

	cards[card.ID] = card

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

// STEP 2
// GET for Scryfall request, with parameters from POST (createCard())
func getCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	card, found := cards[id]

	w.Header().Set("Content-Type", "application/json")
	if found {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(card)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{
			"message": fmt.Sprintf("Card with id %s not found", id),
		}
		json.NewEncoder(w).Encode(response)
	}

}

// STEP 3
// Filter received JSON data (name + image uri)

// STEP 4
// GET for Vue App to display the card into the interface
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", serverStatus)
	mux.HandleFunc("GET /cards", getCards)
	mux.HandleFunc("GET /cards/{id}", getCard)
	mux.HandleFunc("POST /cards", createCard)
	// mux.HandleFunc("DELETE /cards/{id}", deleteCard)
	// mux.HandleFunc("PUT /cards/{id}", updateCard)

	mux.HandleFunc("GET /", serverStatus)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
