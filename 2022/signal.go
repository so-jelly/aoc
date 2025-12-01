package aoc2022

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CalcSignals(b []byte) (int, string) {
	xRegister := 1
	var cycles, signalSum, currentPos int
	nextCycle := 20
	var draw string

	const (
		lit   = "#"
		unlit = "."
	)
	var litOrUnlit string

	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var addX int
		var err error
		addCycle := 1
		if line[0] == "addx" {
			addCycle = 2
			addX, err = strconv.Atoi(line[1])
			if err != nil {
				panic(err)
			}
		}
		if cycles+addCycle >= nextCycle {
			signalSum = signalSum + nextCycle*xRegister
			nextCycle = nextCycle + 40
		}
		for i := addCycle; i > 0; i-- {
			crtLines := []int{1, 2, 3, 4, 5, 6}
			if cycles%40 == 0 && slices.Contains(crtLines, cycles/40) {
				draw = draw + "\n"
				currentPos = 0
			}
			litOrUnlit = unlit
			registerToPosition := Abs(currentPos - xRegister)
			if registerToPosition < 2 {
				litOrUnlit = lit
			}
			draw = draw + fmt.Sprint(litOrUnlit)
			cycles++
			currentPos++
		}

		if cycles > 242 {
			break
		}
		xRegister = xRegister + addX
	}
	return signalSum, draw
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
