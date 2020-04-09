package main

import "fmt"

func main() {
	b := initGameBoard()
	b.player1 = humanPlayer{playerNumber: 1}
	b.player2 = aiPlayer{playerNumber: -1}

	play(b)

}

func play(g gameBoard) {
	gameOver := false
	winner := 0
	g.activePlayerNumber = 1

	for !gameOver {
		if g.movesPlayed%2 == 0 {
			g.player1.makeMove(g)
		} else {
			g.player2.makeMove(g)
		}
		g.movesPlayed++
		gameOver, winner = g.checkGameOver()
		g.displayBoard()
		fmt.Println()
		g.activePlayerNumber *= -1
	}
	if winner == 1 {
		fmt.Println("Winner: X")
	} else if winner == -1 {
		fmt.Println("Winner: O")
	} else {
		fmt.Println("Draw")
	}

}
