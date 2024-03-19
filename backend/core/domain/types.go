package domain

import "github.com/google/uuid"

type Suit struct {
	Index uint8
	Name  string
	Code  string
}

type CardType struct {
	Index uint8
	Name  string
	Code  string
}

type Card struct {
	Value       string `json:"value"`
	Suit        string `json:"suit"`
	Code        string `json:"code"`
	GlobalIndex int
}

type Deck struct {
	Id       uuid.UUID `json:"id"`
	Shuffled bool      `json:"shuffled"`
	Cards    []Card    `json:"cards"`
}
