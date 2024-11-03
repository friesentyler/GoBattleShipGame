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
