package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pid int
var opp_id int
var gameBoard Board
var validBoards []int
var timebank int
var time_per_move int

func getInt(strVal string) int {
	val, err := strconv.Atoi(strVal)
	if err != nil {
		panic(err)
	}
	return val
}

func storeSettings(settings []string) {
	switch field := settings[0]; field {
	case "timebank":
		timebank = getInt(settings[1])
	case "time_per_move":
		time_per_move = getInt(settings[1])
	case "your_botid":
		pid = getInt(settings[1])
		opp_id = 3 - pid
	}
}

func updateGame(updates []string) {
	switch update := updates[0]; update {
	case "field":
		gameBoard = projectFields(updates[1])
	case "macroboard":
		validBoards = make([]int, 0, 9)
		boardStrings := strings.Split(updates[1], ",")
		for i, boardString := range boardStrings {
			if isPlayable := getInt(boardString) == -1; isPlayable {
				validBoards = append(validBoards, i)
			}
		}
	}
}

func printAction(params []string) {
	move := RunMonteCarlo(validBoards, &gameBoard)
	moveString := projectMove(move)
	fmt.Println(moveString)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		switch words := strings.Fields(line); words[0] {
		case "settings":
			storeSettings(words[1:])
		case "update":
			updateGame(words[2:])
		case "action":
			printAction(words[1:])
		default:
			//Something wrong
			fmt.Println("default: " + words[0])
		}
	}
}
