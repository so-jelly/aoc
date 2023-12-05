package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day4() {
	processData(bufio.NewReader(os.Stdin))
	for _, c := range cards {
		c.FindWinningNumbers()
		c.CalculatePoints()
	}

	p := Part()
	switch p {
	case 1:
		var points int
		for _, c := range cards {
			points += c.points
		}
		fmt.Println(points)
	case 2:
		fmt.Println("Part 2")
		fmt.Println(len(cards))
		part2()
	}
}

type Card struct {
	number         int
	drawnNumbers   []int
	cardNumbers    []int
	winningNumbers []int
	wins           int
	points         int
	copies         int
}

var lazySquares map[int]int

// get points.. lazily
func (c *Card) CalculatePoints() {
	if lazySquares == nil {
		lazySquares = map[int]int{0: 0, 1: 1, 2: 2, 3: 4}
	}
	if val, ok := lazySquares[c.wins]; ok {
		c.points = val
		return
	}
	lazySquares[c.wins] = 1 << (c.wins - 1)
	c.points = lazySquares[c.wins]
}

func stringToIntSlice(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	parts := strings.Fields(s)
	nums := make([]int, len(parts))
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" { // some weird spacing in the data
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: '%s' '%s'", part, s)
		}
		nums[i] = num
	}
	return nums, nil
}

func NewCard(number int) *Card { return &Card{number: number, copies: 1} }

func (c *Card) FindWinningNumbers() {
	for _, num := range c.cardNumbers {
		if slices.Contains(c.drawnNumbers, num) {
			c.winningNumbers = append(c.winningNumbers, num)
			c.wins++
		}
	}
}

var cards []*Card

func processData(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		cardLine := scanner.Text()
		colon := strings.Split(cardLine, ":")
		cardStr := strings.TrimPrefix(colon[0], "Card ")
		cardStr = strings.TrimSpace(cardStr)
		cardNum, err := strconv.Atoi(cardStr)
		if err != nil {
			fmt.Println(err)
		}
		c := NewCard(cardNum)
		dataPipe := strings.Split(colon[1], "|")
		c.drawnNumbers, err = stringToIntSlice(dataPipe[0])
		if err != nil {
			panic(err)
		}
		c.cardNumbers, err = stringToIntSlice(dataPipe[1])
		if err != nil {
			panic(err)
		}
		cards = append(cards, c)
	}
}

func part2() {
	var cardCopies int
	for cardIdx, card := range cards {
		for copies := 1; copies <= card.copies; copies++ {
			for win := 1; win <= card.wins; win++ {
				if len(cards) > cardIdx+win {
					cards[cardIdx+win].copies++
				}
			}
		}
		cardCopies += card.copies
	}
	fmt.Println(cardCopies)
}
