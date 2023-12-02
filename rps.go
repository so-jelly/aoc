package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	RockCall         string = "A"
	PaperCall               = "B"
	ScissorsCall            = "C"
	RockResponse            = "X"
	PaperResponse           = "Y"
	ScissorsResponse        = "Z"
)

const (
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	RockValue     int = 1
	PaperValue        = 2
	ScissorsValue     = 3
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	RPSwinValue  = 6
	RPSdrawValue = 3
	RPSloseValue = 0
)

type RPSgame []string

func GetRPSgame(s string) RPSgame {
	game := strings.Fields(s)
	rpsgame := RPSgame(game)
	rpsgame.isValidGame()
	return rpsgame
}

func (g RPSgame) isValidGame() {
	if len(g) != 2 {
		fmt.Println("Invalid game found")
		fmt.Println(g)
		os.Exit(1)
	}

	switch {
	case !slices.Contains([]string{RockCall, PaperCall, ScissorsCall}, (g)[0]):
		fmt.Println("invalid call found")
		os.Exit(1)
	case !slices.Contains([]string{RockResponse, PaperResponse, ScissorsResponse}, (g)[1]):
		fmt.Println("Invalid response found")
		os.Exit(1)
	}

}

func (g RPSgame) Score() int {
	var score int
	call := g[0]
	response := g[1]

	switch {
	// Win
	case (call == RockCall && response == PaperResponse) ||
		(call == PaperCall && response == ScissorsResponse) ||
		(call == ScissorsCall && response == RockResponse):
		score += RPSwinValue
		// Draw
	case (call == RockCall && response == RockResponse) ||
		(call == PaperCall && response == PaperResponse) ||
		(call == ScissorsCall && response == ScissorsResponse):
		score += RPSdrawValue
		// Lose - adding 0 in case lose value changes
	case (call == RockCall && response == ScissorsResponse) ||
		(call == PaperCall && response == RockResponse) ||
		(call == ScissorsCall && response == PaperResponse):
		score += RPSloseValue
	}

	switch {
	case response == RockResponse:
		score += RockValue
	case response == PaperResponse:
		score += PaperValue
	case response == ScissorsResponse:
		score += ScissorsValue
	}
	return score
}

func ScoreRPStournament(b []byte) (score, optimal int) {
	var total int
	var optimalTotal int
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		rpsgame := GetRPSgame(scanner.Text())
		total += rpsgame.Score()
		optimalTotal += rpsgame.ScoreOptimal()
	}
	return total, optimalTotal
}

const (
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
	RPSoptimalWin  = "Z"
	RPSoptimalDraw = "Y"
	RPSoptimalLose = "X"
)

func (g RPSgame) ScoreOptimal() int {
	var score int
	call := g[0]
	outcome := g[1]

	switch {
	// Win
	case outcome == RPSoptimalWin:
		score += RPSwinValue
	case outcome == RPSoptimalDraw:
		score += RPSdrawValue
	case outcome == RPSoptimalLose:
		score += RPSloseValue
	}

	switch {
	case (call == RockCall && outcome == RPSoptimalWin) ||
		(call == PaperCall && outcome == RPSoptimalDraw) ||
		(call == ScissorsCall && outcome == RPSoptimalLose):
		score += PaperValue
	case (call == RockCall && outcome == RPSoptimalLose) ||
		(call == PaperCall && outcome == RPSoptimalWin) ||
		(call == ScissorsCall && outcome == RPSoptimalDraw):
		score += ScissorsValue
	case (call == RockCall && outcome == RPSoptimalDraw) ||
		(call == PaperCall && outcome == RPSoptimalLose) ||
		(call == ScissorsCall && outcome == RPSoptimalWin):
		score += RockValue
	}
	return score
}
