package battleship

import (
	"math/rand"
	"time"
)

const BoardSize = 10

var ShipSizes = map[string]int {
	"Carrier": 5,
	"Battleship": 4,
	"Cruiser": 3,
	"Submarine": 3,
	"Destroyer": 2,
}

type Board [BoardSize][BoardSize]string

type Game struct {
	PlayerBoard Board
	ComputerBoard Board
	HiddenComputerBoard Board
	PlayerHits int
	ComputerHits int
}

// Function to randomly place ships on the board
func PlaceShips(board *Board) {
	rand.Seed(time.Now().UnixNano())

	for name, size := range ShipSizes {
		placed := false
		for !placed {
			orientation := rand.Intn(2) // 0 is horizontal and 1 is vertical
			row, col := rand.Intn(BoardSize), rand.Intn(BoardSize)

			if canPlaceShip(board, row, col, size, orientation) {
				placeShip(board, row, col, size, orientation, name)
				placed = true
			}
		}
	}
}

func canPlaceShip(board *Board, row, col, size, orientation int) bool {
	if orientation == 0 {
		if col+size > BoardSize {
			return false
		}
		for i := 0; i < size; i++ {
			if board[row][col+i] != "" {
				return false
			}
		}
	} else {
		if row+size > BoardSize {
			return false
		}
		for i := 0; i < size; i++ {
			if board[row+i][col] != "" {
				return false
			}
		}
	}
	return true
}

func placeShip(board *Board, row, col, size, orientation int, name string) {
	if orientation == 0 {
		for i := 0; i < size; i++ {
			board[row][col+i] = name
		}
	} else {
		for i := 0; i < size; i++ {
			board[row+i][col] = name
		}
	}
}

func (g *Game) CheckMove(board *Board, row int, col int) string {
	if board[row][col] != "" && board[row][col] != "X" && board[row][col] != "O" {
		shipName := board[row][col]
		board[row][col] = "X"
		g.PlayerHits++
		return "Hit " + shipName + "!"
	}
	board[row][col] = "O"
	return "Miss!"
}

func (g *Game) PlayerMove(row int, col int) string {
	result := g.CheckMove(&g.ComputerBoard, row, col)
	g.GenerateHiddenBoard()
	return result
}

func (g *Game) GenerateHiddenBoard() {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.ComputerBoard[row][col] == "X" {
				g.HiddenComputerBoard[row][col] = "X"
			} else if g.ComputerBoard[row][col] == "O" {
				g.HiddenComputerBoard[row][col] = "O"
			} else {
				g.HiddenComputerBoard[row][col] = ""
			}
		}
	}
}

func (g *Game) IsWon() bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.ComputerBoard[row][col] != "" && g.ComputerBoard[row][col] != "X" && g.ComputerBoard[row][col] != "O" {
				return false
			}
		}
	}
	return true
}

func (g *Game) IsLost() bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.PlayerBoard[row][col] != "" && g.PlayerBoard[row][col] != "X" && g.PlayerBoard[row][col] != "O" {
				return false
			}
		}
	}
	return true
}

func (g *Game) ComputerMove() (int, int, string) {
	rand.Seed(time.Now().UnixNano())
	for {
		row := rand.Intn(BoardSize)
		col := rand.Intn(BoardSize)

		if g.PlayerBoard[row][col] == "" || g.PlayerBoard[row][col] != "X" && g.PlayerBoard[row][col] != "O" {
			result := g.CheckMove(&g.PlayerBoard, row, col)
			return row, col, result
		}
	}
}
