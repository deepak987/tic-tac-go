package main

import "fmt"

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
	player2 ticTacToePlayer // Or rather, player -1
}

func (g gameBoard) play() {
	gameOver := false
	winner := 0
	movesPlayed := 0
	for !gameOver {
		if movesPlayed%2 == 0 {
			g.player1.makeMove(g)
		} else {
			g.player2.makeMove(g)
		}
		movesPlayed++
		gameOver, winner = g.checkGameOver()
		g.displayBoard()
		fmt.Println()
	}
	fmt.Printf("Winner: %d", winner)
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

func (g gameBoard) displayBoard() {
	for i := 0; i < BoardDimension; i++ {
		for j := 0; j < BoardDimension; j++ {
			if g.board[i][j] == 1 {
				fmt.Print("X")
			} else if g.board[i][j] == -1 {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Print("\n")
	}
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

func (g gameBoard) checkGameOver() (bool, int) {
	//Check Rows
	for i := 0; i < BoardDimension; i++ {
		allSame := true
		for j := 1; j < BoardDimension; j++ {
			if g.board[i][j] != g.board[i][j-1] {
				allSame = false

			}

		}
		if allSame {
			if g.board[i][0] != 0 {
				return true, g.board[i][0]
			}
		}
	}

	//Columns
	for j := 0; j < BoardDimension; j++ {
		allSame := true
		for i := 1; i < BoardDimension; i++ {
			if g.board[i][j] != g.board[i-1][j] {
				allSame = false

			}

		}
		if allSame {
			if g.board[0][j] != 0 {
				return true, g.board[0][j]
			}
		}
	}
	// Diagonal 1
	allSame := true
	for i := 1; i < BoardDimension; i++ {
		if g.board[i-1][i-1] != g.board[i][i] {
			allSame = false

		}
	}
	if allSame {
		if g.board[0][0] != 0 {
			return true, g.board[0][0]
		}
	}
	// Diagonal 2
	allSame = true
	for i := 1; i < BoardDimension; i++ {
		j := BoardDimension - i
		if g.board[i][j-1] != g.board[i-1][j] {
			allSame = false

		}
	}
	if allSame {
		if g.board[0][BoardDimension-1] != 0 {
			return true, g.board[0][BoardDimension-1]
		}
	}

	if len(g.getAllEmptyLocations()) == 0 {
		//draw
		return true, 0

	}
	return false, 0

}
