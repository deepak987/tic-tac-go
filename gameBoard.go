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
	board   [][]int
	player1 ticTacToePlayer
	player2 ticTacToePlayer
}

func (g gameBoard) checkGameOver() int {
	return 0
}

func (g gameBoard) getAllEmptyLocations() []boardCoordinate {
	var emptyLocations []boardCoordinate
	for i := 0; i < BoardDimension; i++ {
		for j := 0; j < BoardDimension; j++ {
			if g.board[i][j] == 0 {
				emptyLocations = append(emptyLocations, boardCoordinate{i, j})
			}
		}
	}
	return emptyLocations
}

func initGameBoard() gameBoard {
	b := gameBoard{}
	a := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	b.board = a
	return b
}
