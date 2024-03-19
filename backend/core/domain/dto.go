package domain

import "github.com/google/uuid"

type CreateDeckDto struct {
	Shuffled bool   `json:"shuffled"`
	Cards    string `json:"cards"`
}

type DrawCardDto struct {
	Id    uuid.UUID `json:"id"`
	Count int       `json:"count"`
}

type PersonDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string
	Phone string
	TaxId string
	Email string
}
