package persistence

import (
	"github.com/google/uuid"
)

type DeckDbRow struct {
	Id       uuid.UUID `db:"id"`
	Cards    string    `db:"cards"`
	Shuffled bool      `db:"shuffled"`
}

type IDS interface {
	string | int | uuid.UUID
}

type IRepository[ID IDS, T any] interface {
	Open(id ID) (*T, error)
	Create(payload T) (*T, error)
}
