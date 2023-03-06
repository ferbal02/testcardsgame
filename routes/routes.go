// Author: Ferran Balaguer

package routes

import (
	"net/http"
	"test/cardsgame/api"
	"test/cardsgame/controllers"
	"test/cardsgame/data"

	"github.com/gin-gonic/gin"
)

// Initialise all required routes and handlers
func InitialiseRoutes() *gin.Engine {

	router := gin.Default()

	// Welcome page
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is the Cards Game API")
	})

	// Serves swagger-ui page as static content
	router.Static("/swagger/v1", "./dist")

	// REST Handler Initialisation
	deckRepo := &data.MemoryDeckRepository{}
	deckController := controllers.NewDeckController(deckRepo)
	deckHandler := api.NewDeckHandler(deckController)

	// REST Routes definition

	api := router.Group("/api/v1")
	api.POST("/deck", deckHandler.CreateDeck)
	api.GET("/deck/:uuid", deckHandler.OpenDeck)
	api.GET("/deck/:uuid/cards", deckHandler.DrawCard)

	return router
}
