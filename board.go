package main

import (
	"math"
	"math/rand"
)

type Board [81]*bool

var windices = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{2, 4, 6},
	{0, 4, 8}}

func (board *Board) GetValidMoves(validBoardIndices []int) []int {
	validBoardMoves := make([]int, 0, 81)
	for _, boardIndex := range validBoardIndices {
		offset := boardIndex * 9
		for i := offset; i < offset+9; i++ {
			if board[i] == nil {
				validBoardMoves = append(validBoardMoves, i)
			}
		}
	}
	return validBoardMoves
}

func getBoardIndexOfLast(lastMove int) int {
	return int(math.Mod(float64(lastMove), 9.0))
}

func (board *Board) IsSubBoardPlayable(subBoardIndex int) bool {
	offset := subBoardIndex * 9
	subBoard := board[offset : offset+9]

	//Check if a winner exists in subboard
	for i := 0; i < len(windices); i++ {
		windex := windices[i]
		x := subBoard[windex[0]]
		y := subBoard[windex[1]]
		z := subBoard[windex[2]]
		if x != nil && y != nil && z != nil {
			if *x == *y && *x == *z {
				return false
			}
		}
	}

	//Check if an empty space exists
	for i := 0; i < 9; i++ {
		if subBoard[i] == nil {
			return true
		}
	}
	return false
}

func (board *Board) HasValidMoves(lastMove int) bool {
	lastBoardIndex := getBoardIndexOfLast(lastMove)
	if board.IsSubBoardPlayable(lastBoardIndex) {
		return true
	}
	for subBoard := 0; subBoard < 9; subBoard++ {
		if board.IsSubBoardPlayable(subBoard) {
			return true
		}
	}
	return false
}

func (board *Board) GetSubBoardMoves(boardIndex int) []int {
	moves := make([]int, 0, 9)
	offset := boardIndex * 9
	for move := offset; move < offset+9; move++ {
		moves = append(moves, move)
	}
	return moves
}

func (board *Board) GetAllMoves(lastMove int) []int {
	lastBoardIndex := getBoardIndexOfLast(lastMove)
	if board.IsSubBoardPlayable(lastBoardIndex) {
		return board.GetSubBoardMoves(lastBoardIndex)
	}
	moves := make([]int, 0, 81)
	for subBoard := 0; subBoard < 9; subBoard++ {
		if board.IsSubBoardPlayable(subBoard) {
			subBoardMoves := board.GetSubBoardMoves(subBoard)
			moves = append(moves, subBoardMoves...)
		}
	}
	return moves
}

func (board *Board) GetRandomMove(lastMove int) int {
	moves := board.GetAllMoves(lastMove)
	index := rand.Intn(len(moves))
	return moves[index]
}

func (board *Board) ApplyMove(move int, player bool) {
	b := player
	board[move] = &b
}

func (board *Board) GetWinner() *bool {
	macroboard := make([]*bool, 9, 9)
	for subBoardIndex := 0; subBoardIndex < 9; subBoardIndex++ {
		offset := subBoardIndex * 9
		subBoard := board[offset : offset+9]

		for i := 0; i < len(windices); i++ {
			windex := windices[i]
			x := subBoard[windex[0]]
			y := subBoard[windex[1]]
			z := subBoard[windex[2]]
			if x != nil && y != nil && z != nil {
				if *x == *y && *x == *z {
					macroboard[subBoardIndex] = x
				}
			}
		}
	}
	for i := 0; i < len(windices); i++ {
		windex := windices[i]
		x := macroboard[windex[0]]
		y := macroboard[windex[1]]
		z := macroboard[windex[2]]
		if x != nil && y != nil && z != nil {
			if *x == *y && *x == *z {
				return x
			}
		}
	}
	return nil
}

func (board *Board) Copy() *Board {
	original := *board
	copy := original
	return &copy
}
