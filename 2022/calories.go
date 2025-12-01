package aoc2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"slices"
)

type ElfCalories []int

// GetElfCalories sums the calories each elf carries
func GetElfCalories(b []byte) *ElfCalories {
	e := &ElfCalories{}
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	var currentElfCalories int

	for scanner.Scan() {
		if scanner.Text() == "" {
			*e = append(*e, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("input should be a list of integers")
			panic(err)
		}
		currentElfCalories += calories
	}
	// add the last elf
	*e = append(*e, currentElfCalories)
	return e
}

// Day one Advent of Code 2022

// MaxCalories gets the maximum calories carried by a single elf
func (e *ElfCalories) MaxCalories() int {
	var maxCalories int
	for _, calories := range *e {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	return maxCalories
}

// SumTopThree gets the calories carried by the top three elves
func (e *ElfCalories) SumTopThree() int {
	var sum int
	topThree := make([]int, 3)
	for _, calories := range *e {
		if calories > topThree[0] {
			topThree[0] = calories
		}
		slices.Sort(topThree)
	}
	for _, calories := range topThree {
		sum += calories
	}
	return sum
}
