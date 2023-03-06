// Author: Ferran Balaguer

package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"test/cardsgame/controllers"
	"test/cardsgame/data"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeckHandler struct {
	controller *controllers.DeckController
}

// Converts a data.Card slice into a CardDto slice
func convertCardSlice(cards []data.Card) []CardDto {

	dtoSlice := make([]CardDto, len(cards))
	for i, v := range cards {
		dtoSlice[i] = *convertCardToCardDto(&v)
	}

	return dtoSlice
}

// Converts a data.Card object to a CardDto object
func convertCardToCardDto(card *data.Card) *CardDto {

	dto := &CardDto{
		Code:  card.Code,
		Value: card.Value.String(),
		Suit:  card.Suit.String(),
	}

	return dto
}

// Mounts deck DTO (with no cards) from deck model class
func convertDeckToDeckNoCardsDto(deck *data.Deck) *DeckNoCardsDto {

	dto := &DeckNoCardsDto{
		Id:        deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}

	return dto
}

// Mounts deck DTO (with cards) from deck model class
func convertDeckToDeckDto(deck *data.Deck) *DeckDto {

	dto := &DeckDto{
		Id:        deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}

	dto.Cards = convertCardSlice(deck.Cards)

	return dto
}

// Constructor injects DeckController dependency
func NewDeckHandler(controller *controllers.DeckController) *DeckHandler {

	handler := &DeckHandler{
		controller: controller,
	}

	return handler
}

// REST handler to create a new deck
func (h *DeckHandler) CreateDeck(c *gin.Context) {

	shuffle := false

	// Read suffle value as bool
	if strings.ToLower(c.Query("shuffle")) == "true" {
		shuffle = true
	}

	// Read cards array and extract codes from paramter
	// Converts input to Upper case so that is case-proof
	codes := strings.Split(strings.ToUpper(c.Query("cards")), ",")
	// if split returns 1 empty element is not valid so
	// we remove it
	if len(codes) == 1 && codes[0] == "" {
		codes = nil
	}

	deck, err := h.controller.CreateDeck(shuffle, codes)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// Mounts the DTO from the model object
	dto := convertDeckToDeckNoCardsDto(deck)

	c.IndentedJSON(http.StatusCreated, dto)
}

// REST handler to retrieve one of the existing decks
func (h *DeckHandler) OpenDeck(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("uuid"))
	// Bad request invalid parameter
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	deck, err := h.controller.OpenDeck(uuid)

	if err != nil {
		if errors.Is(err, controllers.ErrDeckNotFound) {
			c.IndentedJSON(http.StatusNotFound, nil)
			return
		} else {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
	}

	// Mounts the DTO from the model object
	dto := convertDeckToDeckDto(deck)

	c.IndentedJSON(http.StatusOK, dto)
}

// REST handler to get one sigle card from a deck
func (h *DeckHandler) DrawCard(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("uuid"))
	// Bad request invalid parameter
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// Initialises amount with the default value of 1. If there
	// is error parsing the amount number returns error
	amount := 1
	if c.Query("amount") != "" {
		if value, err := strconv.Atoi(c.Query("amount")); err == nil {
			amount = value
		} else {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
	}

	cards, err := h.controller.DrawCards(uuid, amount)

	if err != nil {
		if errors.Is(err, controllers.ErrDeckNotFound) {
			c.IndentedJSON(http.StatusNotFound, nil)
			return
		} else if errors.Is(err, controllers.ErrNotEnoughCards) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		} else {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
	}

	// Mounts the DTO from the model object
	dto := convertCardSlice(cards)

	c.IndentedJSON(http.StatusOK, dto)
}
