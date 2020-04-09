package main

type ticTacToePlayer interface {
	chooseMove() // Return a co-ordinate for
}
type randomPlayer struct {
	playerNumber int
}

type humanPlayer struct {
	playerNumber int
}

func (r randomPlayer) chooseMove(b gameBoard) {

}

func (h humanPlayer) chooseMove(b gameBoard) {

}
