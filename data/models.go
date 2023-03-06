// Author: Ferran Balaguer

// This models definition is different from the api/models.go
// as models in the persistence layer are not usually the
// same as those used by the API

package data

import (
	"github.com/google/uuid"
)

// Constants
const MaxCards int = 52

// CardValue Enum definition
type CardValue int

const (
	Ace CardValue = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Jack
	Queen
	King
)

// Translates each enum value to its string representation
func (v CardValue) String() string {

	switch v {
	case Ace:
		return "ACE"
	case One:
		return "1"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Jack:
		return "JACK"
	case Queen:
		return "QUEEN"
	case King:
		return "KING"
	}
	return "Unknown"
}

// CardSuit enum definition
type CardSuit int

const (
	Spades CardSuit = iota
	Diamonds
	Clubs
	Hearts
)

// Translates each enum value to its string representation
func (s CardSuit) String() string {

	switch s {
	case Spades:
		return "SPADES"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	case Hearts:
		return "HEARTS"
	}
	return "Unknown"
}

// Card type definition
type Card struct {
	Value CardValue
	Suit  CardSuit
	Code  string
}

// Deck type definition
type Deck struct {
	Id        uuid.UUID
	Shuffled  bool
	Remaining int
	Cards     []Card
}
