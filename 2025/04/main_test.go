package main

import (
	"strings"
	"testing"
)

const TestInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

const part1Expected = 13
const part2Expected = 43

func TestDay4(t *testing.T) {
	part1Result := day4(strings.NewReader(TestInput), 1)
	if part1Result != part1Expected {
		t.Errorf("Part 1: expected %d, got %d", part1Expected, part1Result)
	}

	part2Result := day4(strings.NewReader(TestInput), 2)
	if part2Result != part2Expected {
		t.Errorf("Part 2: expected %d, got %d", part2Expected, part2Result)
	}
}

func BenchmarkDay4(b *testing.B) {
	for b.Loop() {
		day4(strings.NewReader(TestInput), 1)
		day4(strings.NewReader(TestInput), 2)
	}
}
