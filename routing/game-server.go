package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-test-server/battleship"
)

func main() {
	r := gin.Default()

	var game battleship.Game

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/index.html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

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