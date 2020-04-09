package main

import (
	"math/rand"
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

}
