// Author: Ferran Balaguer

package controllers_test

import (
	"errors"
	"test/cardsgame/controllers"
	"test/cardsgame/data"
	"testing"

	"github.com/google/uuid"
)

func TestDrawCardIncorrectDeck(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	controller.CreateDeck(false, nil)

	_, err := controller.DrawCards(uuid.New(), 1)

	if err == nil {
		t.Errorf("There should be an error")
	}

	if err != nil && !errors.Is(err, controllers.ErrDeckNotFound) {
		t.Errorf("There should be an error of type %v", controllers.ErrDeckNotFound)
	}
}

func TestDrawCardInvalidAmount(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	deck, _ := controller.CreateDeck(false, nil)

	_, err2 := controller.DrawCards(deck.Id, 0)

	if err2 == nil {
		t.Errorf("There should be an error")
	}

	if err2 != nil && !errors.Is(err2, controllers.ErrInvalidAmount) {
		t.Errorf("There should be an error of type %v", controllers.ErrInvalidAmount)
	}
}

func TestDrawCardNotEnoughCardsLeft(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	deck, _ := controller.CreateDeck(false, nil)

	_, err2 := controller.DrawCards(deck.Id, data.MaxCards+1)

	if err2 == nil {
		t.Errorf("There should be an error")
	}

	if err2 != nil && !errors.Is(err2, controllers.ErrNotEnoughCards) {
		t.Errorf("There should be an error of type %v", controllers.ErrNotEnoughCards)
	}
}

func TestDrawCardOkWithAmount(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	amount := 5
	deck, _ := controller.CreateDeck(false, nil)

	cards, err := controller.DrawCards(deck.Id, amount)

	if err != nil {
		t.Errorf("There should not be an error")
	}

	if len(cards) != amount {
		t.Errorf("The number of cards read should be %d", amount)
	}

	deckAfter, _ := controller.OpenDeck(deck.Id)

	if deckAfter.Remaining != data.MaxCards-amount {
		t.Errorf("Remaining cards were not updated correctly")
	}

	if len(deckAfter.Cards) != data.MaxCards-amount {
		t.Errorf("Cards were not removed from the decks card list")
	}

}
