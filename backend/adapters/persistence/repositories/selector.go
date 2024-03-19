package repositories

import (
	"backend/adapters/persistence"
	"github.com/google/uuid"
)

func GetRepository() persistence.IRepository[uuid.UUID, persistence.DeckDbRow] {
	return DecksRepository{}
}
