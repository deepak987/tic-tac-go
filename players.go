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

func (r randomPlayer) makeMove(b gameBoard) {
	moveOptions := b.getAllEmptyLocations()
	s := rand.NewSource(time.Now().Unix())
	rng := rand.New(s) // initialize local pseudorandom generator
	moveIndex := rng.Intn(len(moveOptions))
	x, y := moveOptions[moveIndex].x, moveOptions[moveIndex].y
	b.board[x][y] = r.playerNumber
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
