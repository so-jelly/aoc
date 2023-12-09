package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	DayFunc[6] = Day6
}

func Day6(part int, r io.Reader) {
	fmt.Println("Day 6")
	day6data(r)
	fmt.Println(len(races))

	var possibleWinConditions = 1
	for _, race := range races {
		possibleWinConditions *= race.winConditions()
	}
	fmt.Printf("part1: %d\n", possibleWinConditions)
	p2winConditions := part2race.winConditions()
	fmt.Printf("part2: %d\n", p2winConditions)
}

type Race struct {
	time     int
	distance int
}

func (r Race) winConditions() int {

	var winConds int
	for buttonHeld := 1; buttonHeld < r.time; buttonHeld++ {
		timeleft := r.time - buttonHeld
		speed := buttonHeld
		if speed*timeleft > r.distance {
			winConds++
		}
	}
	fmt.Printf("race %+v has %d win conditions\n", r, winConds)
	return winConds
}

var races []Race
var part2race Race

func day6data(r io.Reader) {
	// 	Time:      7  15   30
	// Distance:  9  40  200
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			line = strings.TrimPrefix(line, "Time:")
			fmt.Println(line)
			raceTimesStr := strings.Fields(line)
			part2timeStr := strings.Join(raceTimesStr, "")
			part2time, err := strconv.Atoi(part2timeStr)
			if err != nil {
				panic(err)
			}
			part2race.time = part2time

			races = make([]Race, len(raceTimesStr))
			for i, s := range raceTimesStr {
				time, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				races[i].time = time
			}

		}
		if strings.HasPrefix(line, "Distance:") {
			line = strings.TrimPrefix(line, "Distance:")
			distStrings := strings.Fields(line)
			p2distStr := strings.Join(distStrings, "")
			p2dist, err := strconv.Atoi(p2distStr)
			if err != nil {
				panic(err)
			}
			part2race.distance = p2dist

			// races = make([]Race, len(dStr))
			for i, distStr := range distStrings {
				distance, err := strconv.Atoi(distStr)
				if err != nil {
					panic(err)
				}
				races[i].distance = distance
			}
		}
	}
}
