package main

import (
	"strings"
	"testing"
)

const part1Expected = 357
const part2Expected = 3121910778619

func TestDay3(t *testing.T) {
	part1Result := day3(strings.NewReader(TestInput), 1)
	if part1Result != part1Expected {
		t.Errorf("Part 1: expected %d, got %d", part1Expected, part1Result)
	}

	part2Result := day3(strings.NewReader(TestInput), 2)
	if part2Result != part2Expected {
		t.Errorf("Part 2: expected %d, got %d", part2Expected, part2Result)
	}
}

func BenchmarkDay3(b *testing.B) {
	for b.Loop() {
		day3(strings.NewReader(TestInput), 1)
		day3(strings.NewReader(TestInput), 2)
	}
}
