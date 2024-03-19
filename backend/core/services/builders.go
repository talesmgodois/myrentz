package services

import (
	"backend/adapters/persistence"
	"backend/core/domain"
)

func buildDeck(d persistence.DeckDbRow) domain.Deck {
	return domain.Deck{
		Shuffled: d.Shuffled,
	}
}
