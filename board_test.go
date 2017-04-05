package main

import (
	"fmt"
	"testing"
)

var emptyBoardString = "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0"

var emptyMacroBoard = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

func TestGetValidMoves(t *testing.T) {
	board := projectFields(emptyBoardString)
	pid = 1
	opp_id = 2
	for _, val := range board {
		if val != nil {
			fmt.Println(*val)
		}
	}
}
