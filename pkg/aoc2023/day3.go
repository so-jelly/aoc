package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func Day3() {
	getPartsAndSymbols(bufio.NewReader(os.Stdin))
	switch Part() {
	case 1:
		Day3Part1()
		return
	case 2:
		Day3Part2()
		return
	}
}

type PartNumber struct {
	Number  int
	Touches map[int]map[int]struct{}
}

type Gear struct {
	Row           int
	Col           int
	AdjacentParts []*PartNumber
}

var parts []*PartNumber

// if character is a symbol, map the symbol position map[x]y
var symbolsPositions map[int]map[int]struct{}

var gears []*Gear

func Day3Part1() {
	var sum int
	fmt.Println("numbers", len(parts))
	for _, v := range parts {
		if v.IsTouching(symbolsPositions) {
			sum += v.Number
		}
	}
	fmt.Println(sum)
}

func getPartsAndSymbols(r io.Reader) {
	scanner := bufio.NewScanner(r)
	// if character is a symbol, map the symbol position map[x]y
	symbolsPositions = make(map[int]map[int]struct{})
	var lineNumber int
	for scanner.Scan() {
		lineNumber++
		row := scanner.Text()
		var partNumber *PartNumber
		var currentNumberString string
		for col, roon := range row {
			isNumber := unicode.IsDigit(roon)
			if !isNumber && currentNumberString != "" {
				partNumber.addNumber(currentNumberString)
				fmt.Println("adding number", partNumber.Number)
				parts = append(parts, partNumber)
				currentNumberString = ""
			}

			if rune(roon) == '.' {
				continue
			}

			if !isNumber {
				// assume a symbol if not a digit or a dot
				if symbolsPositions[col] == nil {
					symbolsPositions[col] = make(map[int]struct{})
				}
				symbolsPositions[col][lineNumber] = struct{}{}
				if rune(roon) == '*' {
					gears = append(gears, &Gear{
						Row: lineNumber,
						Col: col,
					})
				}
			}

			if isNumber {
				// fmt.Println("is number")
				newNumber := currentNumberString == ""
				if newNumber {
					partNumber = new(PartNumber)
				}
				currentNumberString += string(roon)
				fmt.Println(currentNumberString)
				partNumber.addTouches(col, lineNumber, newNumber)
			}
		}
		if currentNumberString != "" {
			partNumber.addNumber(currentNumberString)
			fmt.Println("adding number", partNumber.Number)
			parts = append(parts, partNumber)
			currentNumberString = ""
		}
	}
}

func Day3Part2() {
	var sum int
	for _, g := range gears {
		fmt.Println("gear", g.Col)
		g.GetAdjacentParts()
		sum += g.GetRatio()
	}
	fmt.Println(sum)
}

func (p *PartNumber) addTouches(col, row int, start bool) {
	// fmt.Printf("%v \n %v\n", start, p)
	rows := map[int]struct{}{
		row - 1: {},
		row:     {},
		row + 1: {},
	}
	if start {
		p.Touches = make(map[int]map[int]struct{})
		p.Touches[col-1] = rows
		p.Touches[col] = rows
	}
	p.Touches[col+1] = rows
}

func (p *PartNumber) addNumber(s string) {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // todo dirty
	}
	p.Number = n
}

func (p *PartNumber) IsTouching(positions map[int]map[int]struct{}) bool {
	for col, rows := range p.Touches {
		for row := range rows {
			if _, ok := positions[col][row]; ok {
				fmt.Println(p.Number, " is touching ", col, row)
				return true
			}
		}
	}
	return false
}

func (g *Gear) GetRatio() int {
	if len(g.AdjacentParts) != 2 {
		return 0
	}
	// part 0??
	return g.AdjacentParts[0].Number * g.AdjacentParts[1].Number
}

func (g *Gear) GetAdjacentParts() {
	pos := map[int]map[int]struct{}{g.Col: {g.Row: {}}}
	for _, p := range parts {
		if p.IsTouching(pos) {
			g.AdjacentParts = append(g.AdjacentParts, p)
		}
	}
}
