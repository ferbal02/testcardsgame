// Author: Ferran Balaguer

package controllers_test

import (
	"errors"
	"test/cardsgame/controllers"
	"test/cardsgame/data"
	"testing"

	"github.com/google/uuid"
)

// Tries to retrieve a Deck but there are no decks
func TestOpenDeckNotFoundWhenEmpty(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	_, err := controller.OpenDeck(uuid.New())

	if err == nil {
		t.Errorf("There should be an error")
	}

	if err != nil && !errors.Is(err, controllers.ErrDeckNotFound) {
		t.Errorf("There should be a : %v", err)
	}
}

// Tries to retrieve a Deck but the deck is not found
// amongst the existing decks
func TestOpenDeckNotFoundNotEmpty(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// Inserts Create a new deck with random uuid
	_, err1 := controller.CreateDeck(false, nil)

	if err1 != nil {
		t.Fatalf("Impossible to create deck")
	}

	// Opens another deck
	_, err2 := controller.OpenDeck(uuid.New())

	if err2 == nil {
		t.Errorf("There should be an error")
	}

	if err2 != nil && !errors.Is(err2, controllers.ErrDeckNotFound) {
		t.Errorf("There should be a : %v", err2)
	}
}

// Tests that the deck is found successfully
func TestOpenDeckOk(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// Inserts Create a new deck with random uuid
	created, err1 := controller.CreateDeck(false, nil)

	if err1 != nil {
		t.Fatalf("Impossible to create deck")
	}

	// Opens another deck
	opened, err2 := controller.OpenDeck(created.Id)

	if err2 != nil {
		t.Fatalf("There should not be an error")
	}

	if opened.Id != created.Id {
		t.Errorf("Created and Opened decks are different")
	}
}
