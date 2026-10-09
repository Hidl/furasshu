package main

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"github.com/hidl/furasshu/internal/database"
)

type Card struct {
	ID          int `json:"id"`
	Lesson      int `json:"lesson"`
	Romaji   string `json:"romaji"`
	Hiragana string `json:"hiragana"`
	Katakana string `json:"katakana"`
	Kanji    string `json:"kanji"`
	Notes	 string `json:"notes"`
}

func (cfg *apiConfig) handleAddCard(w http.ResponseWriter, r *http.Request) {
	
	// Fetch from the request body
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the card into the database
	createCardParams := database.CreateCardParams{
		Lesson:   sql.NullInt32{Int32: int32(card.Lesson), Valid: card.Lesson != 0},
		Romaji:   card.Romaji,
		Hiragana: sql.NullString{String: card.Hiragana, Valid: card.Hiragana != ""},
		Katakana: sql.NullString{String: card.Katakana, Valid: card.Katakana != ""},
		Kanji:    sql.NullString{String: card.Kanji, Valid: card.Kanji != ""},
		Notes:    sql.NullString{String: card.Notes, Valid: card.Notes != ""},
	}
	err = cfg.db.CreateCard(r.Context(), createCardParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Card created successfully"))

}