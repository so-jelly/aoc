package main

import (
	"strings"
	"testing"
)

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

const part1Expected = 3
const part2Expected = 6

func TestDay1(t *testing.T) {
	part1Result := day1(strings.NewReader(testInput), 1)
	if part1Result != part1Expected {
		t.Errorf("Part 1: expected %d, got %d", part1Expected, part1Result)
	}

	part2Result := day1(strings.NewReader(testInput), 2)
	if part2Result != part2Expected {
		t.Errorf("Part 2: expected %d, got %d", part2Expected, part2Result)
	}
}

func BenchmarkDay1(b *testing.B) {
	for b.Loop() {
		day1(strings.NewReader(testInput), 1)
		day1(strings.NewReader(testInput), 2)
	}
}
