package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-test-server/battleship"
)

func main() {
	r := gin.Default()

	var game battleship.Game

	r.POST("/start", func(c *gin.Context) {
		game = battleship.Game{}
		battleship.PlaceShips(&game.PlayerBoard)
		battleship.PlaceShips(&game.ComputerBoard)
		c.JSON(http.StatusOK, gin.H {
			"message": "Game Started!",
			"player_board": game.PlayerBoard,
		})
	})

	r.Run(":8080")
}