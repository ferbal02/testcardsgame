// Author: Ferran Balaguer

package controllers

import (
	"errors"
	"math/rand"
	"test/cardsgame/data"
	"time"

	"github.com/google/uuid"
)

// Controller errors
var (
	ErrInvalidCardCode = errors.New("Invalid Card Code")
	ErrDeckNotFound    = errors.New("Deck not found")
	ErrNotEnoughCards  = errors.New("Not enough cards left")
	ErrInvalidAmount   = errors.New("Invalid amount of cards")
	ErrGeneral         = errors.New("General error")
)

// Controller type contains the bussiness logic
type DeckController struct {
	deckRepo data.DeckRepository
}

// Controller constructor injects DeckRepository dependency
func NewDeckController(repository data.DeckRepository) *DeckController {

	controller := &DeckController{
		deckRepo: repository,
	}

	return controller
}

// Returns the card's code from its suit and value
func (c *DeckController) GenerateCardsCode(suit data.CardSuit, value data.CardValue) string {
	return string(suit.String()[0]) + string(value.String()[0])
}

// Generates a cards set without shuffle
func (c *DeckController) GetDefaultCardSet() []data.Card {

	result := make([]data.Card, data.MaxCards)
	index := 0

	for s := data.Spades; s <= data.Hearts; s++ {
		for v := data.Ace; v <= data.King; v++ {

			card := data.Card{
				Value: v,
				Suit:  s,
				Code:  c.GenerateCardsCode(s, v),
			}

			if index < data.MaxCards {
				result[index] = card
				index++
			}
		}
	}

	return result
}

// Generates a random int slice with the desired length
func (c *DeckController) getRandomIntArray(length int) []int {

	a := make([]int, length)

	// Initialise array
	for i := range a {
		a[i] = i
	}

	// Fisher–Yates shuffle
	rand.Seed(time.Now().UnixNano())
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	return a
}

// Generates a randomly shuffled cards set
func (c *DeckController) GetShuffledCardSet() []data.Card {

	// Creates a default cards set
	defaultCards := c.GetDefaultCardSet()

	random := c.getRandomIntArray(len(defaultCards))

	shuffledCards := make([]data.Card, len(defaultCards))

	// Re-arrange cards
	for i, v := range random {
		shuffledCards[i] = defaultCards[v]
	}

	return shuffledCards
}

// Generates as many cards as codes are passed as
// argument
func (c *DeckController) GetCardSetByCodes(codes []string) ([]data.Card, error) {

	defaultCards := c.GetDefaultCardSet()
	result := make([]data.Card, len(codes))
	exists := false

	for ic, vc := range codes {
		exists = false
		for _, vd := range defaultCards {
			if vc == vd.Code {
				result[ic] = vd
				exists = true
				break
			}
		}

		if !exists {
			break
		}
	}

	if !exists {
		return result, ErrInvalidCardCode
	}

	return result, nil
}

// Creates a deck with the desired cards
// If shuffle is true the card set is randomly shuffled
// If codes has value, only the selected cards are used
func (c *DeckController) CreateDeck(shuffled bool, codes []string) (*data.Deck, error) {

	var cardSet []data.Card
	var err error
	doShuffle := shuffled

	if codes != nil && len(codes) > 0 {
		doShuffle = false
		cardSet, err = c.GetCardSetByCodes(codes)
	} else if shuffled {
		cardSet = c.GetShuffledCardSet()
	} else {
		cardSet = c.GetDefaultCardSet()
	}

	if err != nil {
		return nil, err
	}

	deck := data.Deck{
		Id:        uuid.New(),
		Shuffled:  doShuffle,
		Remaining: len(cardSet),
		Cards:     cardSet,
	}

	// Adds the newly create deck to de Repository
	c.deckRepo.Add(deck)

	return &deck, nil
}

// Tries to retrieve a deck from the repository if exists, otherwise
// returns not found error
func (c *DeckController) OpenDeck(uuid uuid.UUID) (*data.Deck, error) {

	deck, err := c.deckRepo.GetDeckById(uuid)

	if err != nil {
		return nil, ErrDeckNotFound
	}

	return deck, nil
}

// Draws amount cards from the deck, removing them from the deck
// and updating the remaining value
func (c *DeckController) DrawCards(uuid uuid.UUID, amount int) ([]data.Card, error) {

	cards, err := c.deckRepo.DrawCardsFromDeck(uuid, amount)

	if err != nil {
		switch err {
		case data.ErrNotFound:
			return nil, ErrDeckNotFound
		case data.ErrInvalidParameters:
			return nil, ErrInvalidAmount
		case data.ErrTruncate:
			return nil, ErrNotEnoughCards
		}
	}

	return cards, nil
}
