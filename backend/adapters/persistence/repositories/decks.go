package repositories

import (
	"backend/adapters/persistence"
	"github.com/google/uuid"
)

type DecksRepository struct {
}

func (d DecksRepository) Open(id uuid.UUID) (*persistence.DeckDbRow, error) {
	db := persistence.GetDbx()

	var deckRow persistence.DeckDbRow

	sql := "select d.id, d.cards, d.shuffled from decks d where d.id = $1"

	// Executing the SQL statement and scanning the returned ID into createdDeck
	err := db.Get(&deckRow, sql, id)
	if err != nil {
		return nil, err
	}

	return &deckRow, nil
}

func (d DecksRepository) Create(deck persistence.DeckDbRow) (*persistence.DeckDbRow, error) {
	db := persistence.GetDbx()

	var createdDeck = persistence.DeckDbRow{}

	sql := "INSERT INTO decks (cards, shuffled) VALUES ($1, $2) RETURNING id"

	// Executing the SQL statement and scanning the returned ID into createdDeck

	err := db.QueryRowx(sql, deck.Cards, deck.Shuffled).Scan(&createdDeck.Id)
	if err != nil {
		return nil, err
	}

	return &createdDeck, nil
}
