package main

const (
	// BoardDimension - Standard tic-tac-toe board size
	BoardDimension int = 3
)

type boardCoordinate struct {
	x int
	y int
}
type gameBoard struct {
	board   [][]uint
	player1 ticTacToePlayer
	player2 ticTacToePlayer
}

func (g gameBoard) checkGameOver() int {
	return 0
}

func (g gameBoard) getAllEmptyLocations() []boardCoordinate {
	return make([]boardCoordinate, 0, 0)
}
