package main

import (
	"strings"
	"testing"
)

const TestInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

const part1Expected = 3
const part2Expected = 14

func TestDay4(t *testing.T) {
	part1Result, part2Result := day5(strings.NewReader(TestInput))
	if part1Result != part1Expected {
		t.Errorf("Part 1: expected %d, got %d", part1Expected, part1Result)
	}
	if part2Result != part2Expected {
		t.Errorf("Part 2: expected %d, got %d", part2Expected, part2Result)
	}
}

func BenchmarkDay5(b *testing.B) {
	for b.Loop() {
		day5(strings.NewReader(TestInput))
	}
}
