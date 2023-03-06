// Author: Ferran Balaguer

package data

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound          = errors.New("Not found")
	ErrInvalidParameters = errors.New("Invalid argument")
	ErrTruncate          = errors.New("Truncated items")
)

// Data abstraction interface to decouple the
// Bussiness layer from the data layer
type DeckRepository interface {

	// Crates a new deck in the repository
	Add(Deck)
	// Gets a deck from the repostory
	GetDeckById(uuid.UUID) (*Deck, error)
	// Gets a concrete card from a deck
	GetDeckCardByCode(uuid.UUID, string) (*Card, error)
	// Get cards from deck
	DrawCardsFromDeck(uuid.UUID, int) ([]Card, error)
}

// Implements DeckRepository using
// a map in memory as storage
type MemoryDeckRepository struct {
	decks map[uuid.UUID]*Deck
}

// DeckRepository interface implementation

func (r *MemoryDeckRepository) Add(deck Deck) {

	if r.decks == nil {
		r.decks = map[uuid.UUID]*Deck{}
	}

	r.decks[deck.Id] = &deck
}

func (r *MemoryDeckRepository) GetDeckById(uuid uuid.UUID) (*Deck, error) {

	deck, ok := r.decks[uuid]

	if !ok {
		return nil, ErrNotFound
	}

	return deck, nil
}

func (r *MemoryDeckRepository) GetDeckCardByCode(uuid uuid.UUID, code string) (*Card, error) {

	var card *Card

	if deck, ok := r.decks[uuid]; ok {
		for _, v := range deck.Cards {
			if v.Code == code {
				card = &v
				break
			}
		}
	}

	if card == nil {
		return nil, ErrNotFound
	}

	return card, nil
}

func (r *MemoryDeckRepository) DrawCardsFromDeck(uuid uuid.UUID, amount int) ([]Card, error) {

	// Invalid amout error
	if amount <= 0 {
		return nil, ErrInvalidParameters
	}

	if deck, ok := r.decks[uuid]; ok {

		// Error not enough cards left
		if amount > deck.Remaining {
			return nil, ErrTruncate
		}

		cards := deck.Cards[:amount]
		// remove "amount" cards form the top of the cards list
		deck.Cards = deck.Cards[amount:]
		// update remaining
		deck.Remaining -= amount

		return cards, nil
	} else {
		return nil, ErrNotFound
	}
}
