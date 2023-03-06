// Author: Ferran Balaguer

package controllers_test

import (
	"errors"
	"test/cardsgame/controllers"
	"test/cardsgame/data"
	"testing"
)

// Tests that the default card set is correctly
// mounted and has the correct lenght
func TestDefaultCardSet(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	defaultCards := controller.GetDefaultCardSet()

	if defaultCards[0].Code != "SA" ||
		defaultCards[13].Code != "DA" ||
		defaultCards[26].Code != "CA" ||
		defaultCards[39].Code != "HA" {
		t.Errorf("Bad Default Card Set order")
	}

	if len(defaultCards) != data.MaxCards {
		t.Errorf("Bad cars set lentgh")
	}
}

// Tests that two different calls to GetShuffledCardSet
// return cards sorted differently.
func TestShuffledCardSet(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	shuffledCards1 := controller.GetShuffledCardSet()
	shuffledCards2 := controller.GetShuffledCardSet()

	if len(shuffledCards1) != len(shuffledCards2) {
		t.Errorf("2 different calls to GetShuffledCardSet must be the same lenght")
	}

	different := false
	for i, v := range shuffledCards1 {
		if shuffledCards2[i].Code != v.Code {
			different = true
			break
		}
	}

	if !different {
		t.Errorf("Two consecutive shuffled card sets must be different")
	}
}

// Checks that the Deck is full with 52 cards and sorted
func TestCreateDeckDefault(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	deck, _ := controller.CreateDeck(false, nil)

	if len(deck.Cards) != data.MaxCards {
		t.Errorf("There should be %d cards.", data.MaxCards)
	}

	if deck.Shuffled {
		t.Errorf("Shuffled deck property must be false")
	}

	if deck.Remaining != data.MaxCards {
		t.Errorf("Remaining cards must be = %d", data.MaxCards)
	}

	// Deck cards set has to be sorted like the default one
	defaultCards := controller.GetDefaultCardSet()
	for i, v := range defaultCards {
		if v != deck.Cards[i] {
			t.Errorf("Cards are not correctly sorted like default")
		}
	}
}

// Checks that the Deck is full but cards order is
// different from the default one
func TestCreateDeckShuffled(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	deck, _ := controller.CreateDeck(true, nil)

	if len(deck.Cards) != data.MaxCards {
		t.Errorf("There should be %d cards.", data.MaxCards)
	}

	if !deck.Shuffled {
		t.Errorf("Shuffled deck property must be true")
	}

	if deck.Remaining != data.MaxCards {
		t.Errorf("Remaining cards must be = %d", data.MaxCards)
	}

	// Deck cards set has to be sorted like the default one
	sorted := true
	defaultCards := controller.GetDefaultCardSet()
	for i, v := range defaultCards {
		if v != deck.Cards[i] {
			sorted = false
			break
		}
	}

	if sorted {
		t.Errorf("The deck must contain an unsorted cards set")
	}
}

// The cards list is passed but card code is wrong so
// that an error is flagged
func TestCreateDeckPartilaWrongCode(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// XX is an invalid code for a card
	codes := []string{"SA", "XX"}
	_, err := controller.CreateDeck(false, codes)

	if err == nil {
		t.Errorf("Should have returned %v", controllers.ErrInvalidCardCode)
	}
}

// Tests that if we pass an nonexistent code in the cards set
// an error is flagged
func TestGetCardSetByCodesErr(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// The second one is wrong
	_, err := controller.GetCardSetByCodes([]string{"SA", "YY"})

	if err == nil || !errors.Is(err, controllers.ErrInvalidCardCode) {
		t.Errorf("Should have returned %v", controllers.ErrInvalidCardCode)
	}
}

// Tests that if all codes in passes as argument exists, a slice
// with the corresponding card is created
func TestGetCardSetByCodesOK(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// The second one is wrong
	cards, err := controller.GetCardSetByCodes([]string{"SK", "H8"})

	if err != nil {
		t.Errorf("There should be no error: %v", err)
	}

	if len(cards) != 2 {
		t.Errorf("Cards set length should be 2")
	}

	if cards[0].Code != "SK" {
		t.Errorf("First card should be SK")
	}

	if cards[1].Code != "H8" {
		t.Errorf("Second card should be H8")
	}
}

// Tests that passing the desired card codes to the
// deck generator returns the wanted card objects
func TestCrateDeckPartialOk(t *testing.T) {

	repository := &data.MemoryDeckRepository{}
	controller := controllers.NewDeckController(repository)

	// XX is an invalid code for a card
	codes := []string{"SQ", "D5"}
	deck, err := controller.CreateDeck(true, codes)

	if err != nil {
		t.Errorf("There should not be any error")
	}

	if len(deck.Cards) != len(codes) {
		t.Errorf("There should be %d cards.", len(codes))
	}

	if deck.Shuffled {
		t.Errorf("Shuffled deck property must be false")
	}

	if deck.Remaining != len(codes) {
		t.Errorf("Remaining cards must be = %d", len(codes))
	}

	if deck.Cards[0].Code != codes[0] {
		t.Errorf("First card should be %v", codes[0])
	}

	if deck.Cards[1].Code != codes[1] {
		t.Errorf("Second card should be %v", codes[1])
	}
}
