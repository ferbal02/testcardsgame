// Author: Ferran Balaguer

// This models definition is different from the data/models.go
// as models in the persistence layer are not usually the
// same as those used by the API

package api

import (
	"github.com/google/uuid"
)

// CardDto type definition
type CardDto struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// DeckDto type definition
type DeckDto struct {
	Id        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []CardDto `json:"cards"`
}

// DeckDto type definition
type DeckNoCardsDto struct {
	Id        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}
