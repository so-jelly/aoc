package aoc2022

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestLettersToInts(t *testing.T) {
	in := "azAZ"
	expect := []int{1, 26, 27, 52}
	out := LettersToInts(in)
	if slices.Compare(out, expect) != 0 {
		t.Errorf("LettersToInts(%v) = %v, want %v", in, out, expect)
	}
}

var dayThreeTestData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

var dayThreeTestResult = 157

func TestBagPriority(t *testing.T) {
	priorityTotal := BagPriorities([]byte(dayThreeTestData))
	if priorityTotal != dayThreeTestResult {
		t.Errorf("FindPriority(%v) = %v, want %v", dayThreeTestData, priorityTotal, dayThreeTestResult)
	}
}
