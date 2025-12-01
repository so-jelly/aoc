package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	NS = "|"
	EW = "-"
	NE = "L"
	NW = "J"
	SW = "7"
	SE = "F"
	S  = "S"
	G  = "."
)

var tileToConnections = map[string]string{NS: "NS", EW: "EW", NE: "NE", NW: "NW", SW: "SW", SE: "SE", S: "NEWS"}

func init() {
	DayFunc[10] = Day10
}

func Day10(part int, r io.Reader) {
	fmt.Println("day10")
	start, rows := d10parse(r)
	fmt.Println(start, rows)
	pipe := FollowPipes(rows, [][][]int{{start}})
	fmt.Println(pipe)
}

func d10parse(r io.Reader) ([]int, [][]string) {
	var s string
	scanner := bufio.NewScanner(r)
	var start []int
	var rows [][]string
	var row int
	for scanner.Scan() {
		line := scanner.Text()
		r := make([]string, len(line))
		for col, c := range line {
			r[col] = string(c)
			if c == 'S' {
				start = append(start, col, row)
			}
		}
		s += scanner.Text() + "\n"
		rows = append(rows, r)
		row++
	}
	return start, rows
}

func FollowPipes(grid [][]string, pipes [][][]int) [][]int {
	var newPipes [][][]int
	for _, pipe := range pipes {
		head := pipe[0]
		tail := pipe[len(pipe)-1]
		tailConnections := InboundConnections(grid, tail)
		if len(tailConnections) == 0 {
			// dead end
			continue
		}
		fmt.Println("found connections:", tailConnections)
		for _, newConnection := range tailConnections {
			if len(pipe) > 4 && // smallest possible pipe
				newConnection[0] == head[0] &&
				newConnection[1] == head[1] {
				// found the head
				return pipe
			}
			if len(pipe) > 2 {
				// can't go back the way we came
				previousConnection := pipe[len(pipe)-2]
				if previousConnection[0] == newConnection[0] &&
					previousConnection[1] == newConnection[1] {
					fmt.Println("skipping", newConnection)
					continue
				}
			}
			newPipes = append(newPipes, append(pipe, newConnection))
		}
	}
	fmt.Println("new pipes:", newPipes)
	return FollowPipes(grid, newPipes)
}

func InboundConnections(grid [][]string, position []int) [][]int {
	// copilot wrote this... looks ok but crossing fingers
	var inboundConnections [][]int
	tile := grid[position[0]][position[1]]
	tileConnections := tileToConnections[tile]
	// from west
	if strings.Contains(tileConnections, "W") && position[0]-1 >= 0 {
		westTile := grid[position[1]][position[0]-1]
		if strings.Contains(tileToConnections[westTile], "E") {
			inboundConnections = append(inboundConnections, []int{position[0] - 1, position[1]})
		}
	}
	// from east
	if strings.Contains(tileConnections, "E") && position[0]+1 < len(grid[0]) {
		eastTile := grid[position[1]][position[0]+1]
		if strings.Contains(tileToConnections[eastTile], "W") {
			inboundConnections = append(inboundConnections, []int{position[0] + 1, position[1]})
		}
	}
	// from north
	if strings.Contains(tileConnections, "N") && position[1]-1 >= 0 {
		northTile := grid[position[1]-1][position[0]]
		if strings.Contains(tileToConnections[northTile], "S") {
			inboundConnections = append(inboundConnections, []int{position[0], position[1] - 1})
		}
	}
	// from south
	if strings.Contains(tileConnections, "S") && position[1]+1 < len(grid) {
		southTile := grid[position[1]+1][position[0]]
		if strings.Contains(tileToConnections[southTile], "N") {
			inboundConnections = append(inboundConnections, []int{position[0], position[1] + 1})
		}
	}
	return inboundConnections
}

// func possiblePipes(s start, rows [][]string) {
// 	var pipes []string
// 	for _, row := range rows {
// 		for _, c := range row {
// 			if c == NS || c == EW {
// 				pipes = append(pipes, c)
// 			}
// 		}
// 	}
// 	fmt.Println(pipes)
// }
