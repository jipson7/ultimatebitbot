package main

import (
	"fmt"
	"math"
	"time"
)

var debug = false

var TIME_PER_MOVE float64 = 623.0 //Time Per Move in Milliseconds

func RunMonteCarlo(validBoards []int, board *Board) int {
	startTime := time.Now()

	movesToTry := board.GetValidMoves(validBoards)

	ties := make(map[int]float64)
	wins := make(map[int]float64)
	weightedWins := make(map[int]float64)
	losses := make(map[int]float64)
	weightedLosses := make(map[int]float64)
	gamesPlayed := 0

	for {

		for _, move := range movesToTry {
			gamesPlayed += 1

			previousPlayer := true

			var movesToGameOver float64

			simulatedBoard := board.Copy()
			simulatedMove := move

			simulatedBoard.ApplyMove(simulatedMove, previousPlayer)
			movesToGameOver += 1.0

			for simulatedBoard.GetWinner() == nil {
				if !simulatedBoard.HasValidMoves(simulatedMove) {
					ties[move] += 1.0
					break
				}

				movesToGameOver += 1.0

				simulatedMove = simulatedBoard.GetRandomMove(simulatedMove)
				previousPlayer = !previousPlayer
				simulatedBoard.ApplyMove(simulatedMove, previousPlayer)
			}
			simulatedWinner := simulatedBoard.GetWinner()
			if simulatedWinner != nil {
				if *simulatedWinner {
					wins[move] += 1.0
					weightedWins[move] += (1.0 / movesToGameOver)
				} else {
					losses[move] += 1.0
					weightedLosses[move] += (1.0 / movesToGameOver)
				}
			}
		}

		end := time.Now()
		if end.Sub(startTime).Seconds()*1000.0 > TIME_PER_MOVE {
			break
		}
	}

	if debug {
		fmt.Printf("%v games played\n", gamesPlayed)
	}

	bestScore := math.Inf(-1)
	var bestMove int

	for _, move := range movesToTry {
		score := (weightedWins[move] - (weightedLosses[move] * 2.0)) / (wins[move] + losses[move] + ties[move])
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	return bestMove
}
