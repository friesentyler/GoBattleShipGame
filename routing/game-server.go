package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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
		game.GenerateHiddenBoard()
		c.JSON(http.StatusOK, gin.H {
			"message": "Game Started!",
			"player_board": game.PlayerBoard,
			"computer_board": game.HiddenComputerBoard,
		})
	})

	r.POST("/move", func(c *gin.Context) {
		rowStr := c.Query("row")
		colStr := c.Query("col")

		row, err1 := strconv.Atoi(rowStr)
		col, err2 := strconv.Atoi(colStr)

		if err1 != nil || err2 != nil || row < 0 || row >= battleship.BoardSize || col < 0 || col >= battleship.BoardSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid row or column"})
			return
		}

		playerResult := game.PlayerMove(row, col)

		compRow, compCol, compResult := game.ComputerMove()

		c.JSON(http.StatusOK, gin.H{
			"player_move": gin.H{
				"row": row,
				"col": col,
				"result": playerResult,
			},
			"computer_move": gin.H{
				"row": compRow,
				"col": compCol,
				"result": compResult,
			},
			"player_board": game.PlayerBoard,
			"computer_board": game.HiddenComputerBoard,
		})

	})

	r.Run(":8080")
}