package main

import (
	"testing"
)

func testgetAllEmptyLocations(t *testing.T) {
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
