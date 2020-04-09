package main

import (
	"fmt"
	"os"
)

func main() {
	b := initGameBoard()
	fmt.Println("Player 1")

	b.player1 = getPlayerFromConsoleInput(1)
	fmt.Println("\n\nPlayer 2")
	b.player2 = getPlayerFromConsoleInput(-1)

	play(b)

}
func getPlayerFromConsoleInput(playerNumber int) ticTacToePlayer {
	fmt.Println("Enter option \n 1. An AI player\n 2. A human player (you)\n 3. A player that always chooses a random square ")
	var x int
	fmt.Scanf("%d", &x)
	if x == 1 {
		return aiPlayer{playerNumber: playerNumber}
	} else if x == 2 {
		return humanPlayer{playerNumber: playerNumber}
	} else if x == 3 {
		return randomPlayer{playerNumber: playerNumber}
	}
	fmt.Println("Illegal choice. Start program again")
	os.Exit(1)
	return nil
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
