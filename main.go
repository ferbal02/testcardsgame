package main

import (
	"test/cardsgame/routes"
)

func main() {

	// Setup and start server
	router := routes.InitialiseRoutes()
	router.Run("localhost:8080")
}
