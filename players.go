package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type ticTacToePlayer interface {
	makeMove(g gameBoard)
}
type randomPlayer struct {
	playerNumber int
}

type humanPlayer struct {
	playerNumber int
}

type aiPlayer struct {
	playerNumber int
}

func (r randomPlayer) makeMove(b gameBoard) {
	moveOptions := b.getAllEmptyLocations()
	s := rand.NewSource(time.Now().Unix())
	rng := rand.New(s) // initialize local pseudorandom generator
	moveIndex := rng.Intn(len(moveOptions))
	x, y := moveOptions[moveIndex].x, moveOptions[moveIndex].y
	b.board[x][y] = r.playerNumber
}

func argmaxSlice(s []int) int {
	maxLoc := 0
	maxVal := s[0]
	for i, v := range s {
		if v > maxVal {
			maxLoc = i
			maxVal = v
		}
	}
	return maxLoc
}
func argminSlice(s []int) int {
	minLoc := 0
	minVal := s[0]
	for i, v := range s {
		if v < minVal {
			minLoc = i
			minVal = v
		}
	}
	return minLoc
}
func (a aiPlayer) makeMove(b gameBoard) {
	var maximizingPlayer bool
	if b.player1 == a {
		maximizingPlayer = true
	} else {
		maximizingPlayer = false
	}
	possibleMoves, possibleBoards := b.getPossibleNextBoards()
	boardEvaluations := make([]int, len(possibleBoards))
	for i, board := range possibleBoards {
		boardEvaluations[i] = minimax(board, !maximizingPlayer)
	}
	var bestMove boardCoordinate
	if maximizingPlayer {
		bestMove = possibleMoves[argmaxSlice(boardEvaluations)]
	} else {
		bestMove = possibleMoves[argminSlice(boardEvaluations)]
	}
	b.board[bestMove.x][bestMove.y] = a.playerNumber

}
func (h humanPlayer) makeMove(b gameBoard) {
	fmt.Println("Enter row number and column number(0 indexed) separated by a space")
	var x int
	var y int
	var symbol string
	legalMoves := b.getAllEmptyLocations()
	if h.playerNumber == 1 {
		symbol = "X"
	} else {
		symbol = "O"
	}

	moveLegal := false
	for !moveLegal {
		fmt.Println(symbol, ": your turn")
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		playerMove := boardCoordinate{x: x, y: y}
		for _, move := range legalMoves {
			if reflect.DeepEqual(playerMove, move) {
				moveLegal = true
				break
			}
		}
		if !moveLegal {
			fmt.Println("Illegal Move")
		}
	}
	b.board[x][y] = h.playerNumber

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minimax(g gameBoard, maximizingPlayer bool) int {
	finished, winner := g.checkGameOver()
	if finished {
		return winner
	}
	var value int
	_, possibleNextBoards := g.getPossibleNextBoards()
	if maximizingPlayer {
		value = -1 // Lowest possible value since no depth limit
		for _, child := range possibleNextBoards {
			value = max(value, minimax(child, false))
		}
	} else {
		value = 1
		for _, child := range possibleNextBoards {
			value = min(value, minimax(child, true))
		}
	}
	return value
}
