package main

func main() {
	r := gin.Default()

	var game Game

	r.POST("/start", func(c *gin.Context) {
		game = Game{}
		placeShips(&game.PlayerBoard)
		placeShips(&game.ComputerBoard)
		c.JSON(http.StatusOK, gin.H {
			"message": "Game Started!",
			"player_board": game.PlayerBoard,
		})
	})

	r.Run(":8080")
}