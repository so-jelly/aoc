package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func init() {
	DayFunc[2] = Day2
}

func Day2(part int, r io.Reader) {
	scanner := bufio.NewScanner(r)
	games := make([]*Game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		game := ParseGame(line)
		game.MaxCubeCount()
		games = append(games, game)
	}
	switch part {
	case 1:
		var sum int
		for _, game := range games {
			if game.PossibleGame() {
				sum += game.Id
			}
		}
		fmt.Println(sum)
		return
	case 2:
		var sum int
		for _, game := range games {
			sum += game.GamePower()
		}
		fmt.Println(sum)
		return
	}

	fmt.Println("not implemented")
}

type Game struct {
	Id        int
	Peeks     []map[string]int
	MaxCounts map[string]int
}

func (g *Game) MaxCubeCount() {
	maxCounts := make(map[string]int)
	for _, peek := range g.Peeks {
		for color, count := range peek {
			if count > maxCounts[color] {
				maxCounts[color] = count
			}
		}
	}
	g.MaxCounts = maxCounts
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

func ParseGame(s string) *Game {
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	g := &Game{}
	gid := strings.Split(s, ":")
	gameIdStr := strings.Split(gid[0], " ")[1]
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		fmt.Printf("invalid game ID: %s \n %v", gameIdStr, err)
		os.Exit(1)
	}
	g.Id = gameId
	p := make([]map[string]int, 0)
	for _, peek := range strings.Split(gid[1], ";") {
		cubes := strings.Split(peek, ",")
		cubeMap := make(map[string]int)
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			c := strings.Split(cube, " ")

			count, err := strconv.Atoi(c[0])
			if err != nil {
				fmt.Printf("invalid cube count: %s \n %v", c[0], err)
			}
			cubeMap[c[1]] = count
		}
		p = append(p, cubeMap)
	}
	g.Peeks = p
	return g
}

func (g *Game) PossibleGame() bool {
	var (
		maxred   = 12
		maxgreen = 13
		maxblue  = 14
	)
	if g.MaxCounts["red"] > maxred ||
		g.MaxCounts["green"] > maxgreen ||
		g.MaxCounts["blue"] > maxblue {
		return false
	}
	return true
}

func (g *Game) GamePower() int {
	power := 1
	for _, c := range g.MaxCounts {
		if c != 0 {
			power = power * c
		}
	}
	return power
}
