package main

import (
	"testing"
)

func TestGetAllEmptyLocations(t *testing.T) {
	b := initGameBoard()
	allEmpty := b.getAllEmptyLocations()
	if len(allEmpty) != BoardDimension*BoardDimension {
		t.Errorf("Expected %d empty, received %d", BoardDimension*BoardDimension, len(allEmpty))
	}
	b.board[0][0] = 1
	allEmpty = b.getAllEmptyLocations()
	if len(allEmpty) != BoardDimension*BoardDimension-1 {
		t.Errorf("Expected %d empty, received %d", BoardDimension*BoardDimension-1, len(allEmpty))
	}
}

func TestGameOverEmpty(t *testing.T) {
	b := initGameBoard()
	finished, _ := b.checkGameOver()
	if finished {
		t.Errorf("Fail- game not finished")
	}
}

func TestGameOverDraw(t *testing.T) {
	b := initGameBoard()
	b.board = [][]int{
		{-1, 1, -1},
		{1, 1, -1},
		{-1, -1, 1},
	}
	finished, winner := b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != 0 {
		t.Errorf("Fail- should be a draw")
	}
}

func TestGameOverRowWin(t *testing.T) {
	b := initGameBoard()
	b.board = [][]int{
		{-1, -1, -1},
		{1, 1, -1},
		{-1, -1, 1},
	}
	finished, winner := b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != -1 {
		t.Errorf("Fail- player 2 should win")
	}

	b.board = [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{-1, -1, 1},
	}
	finished, winner = b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != 1 {
		t.Errorf("Fail- player 2 should win")
	}
}

func TestGameOverColWin(t *testing.T) {
	b := initGameBoard()
	b.board = [][]int{
		{1, -1, -1},
		{1, -1, -1},
		{-1, -1, 1},
	}
	finished, winner := b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != -1 {
		t.Errorf("Fail- player 2 should win")
	}

	b.board = [][]int{
		{1, 1, 1},
		{1, -1, 1},
		{-1, -1, 1},
	}
	finished, winner = b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != 1 {
		t.Errorf("Fail- player 2 should win")
	}
}

func TestGameOverDiagWin(t *testing.T) {
	b := initGameBoard()
	b.board = [][]int{
		{-1, -1, 1},
		{1, -1, -1},
		{-1, 1, -1},
	}
	finished, winner := b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != -1 {
		t.Errorf("Fail- player 2 should win")
	}

	b.board = [][]int{
		{1, 1, 1},
		{-1, 1, -1},
		{1, -1, -1},
	}
	finished, winner = b.checkGameOver()
	if !finished {
		t.Errorf("Fail- game has finished")
	}
	if winner != 1 {
		t.Errorf("Fail- player 2 should win")
	}
}
