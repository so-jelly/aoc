package aoc2022

import (
	"bufio"
	"errors"
	"fmt"
	"slices"
	"strings"
)

var lettersToInt map[string]int

func init() {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	// this is pretty fast
	lettersToInt = map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
		"k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
		"u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30,
		"E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40,
		"O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50,
		"Y": 51, "Z": 52,
	}
}

func LettersToInts(s string) []int {
	runes := []rune(s)
	result := make([]int, len(runes))
	for i := range runes {
		result[i] = lettersToInt[string(runes[i])]
	}
	return result
}

func FindPriority(bag []int) int {
	// Chunk our slice
	a := bag[len(bag)/2:]
	b := bag[:len(bag)/2]
	if len(a) != len(b) {
		panic(errors.New("bag not in half"))
	}
	slices.Sort(a)
	slices.Sort(b)
	for i, n := range a {
		if len(a) > i+1 && a[i+1] == n {
			continue
		}
		if slices.Contains(b, n) {
			return n
		}
	}
	fmt.Println("no match")
	return 0
}

func BagPriorities(b []byte) int {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	var priorityTotal int
	for scanner.Scan() {
		priorityTotal += FindPriority(LettersToInts(scanner.Text()))
	}
	return priorityTotal
}

func GroupPriority(group []string) int {
	a := LettersToInts(group[0])
	slices.Sort(a)
	b := LettersToInts(group[1])
	c := LettersToInts(group[2])

	for i, n := range a {
		if len(a) > i+1 && a[i+1] == n {
			continue
		}
		if slices.Contains(b, n) &&
			slices.Contains(c, n) {
			return n
		}
	}
	fmt.Println("no group priorities")
	return 0
}

func GroupPriorities(b []byte) int {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	var priorityTotal int
	var cursor int
	currentGroup := make([]string, 3)
	for scanner.Scan() {
		currentGroup[cursor] = scanner.Text()
		cursor++
		if cursor == 3 {
			priorityTotal += GroupPriority(currentGroup)
			cursor = 0
		}
	}
	return priorityTotal
}
