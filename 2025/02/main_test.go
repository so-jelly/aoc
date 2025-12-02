package main

import (
	"strings"
	"testing"
)

const TestInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

const part1Expected = 1227775554
const part2Expected = 4174379265

func TestDay2(t *testing.T) {
	part1Result := day2(strings.NewReader(TestInput), 1)
	if part1Result != part1Expected {
		t.Errorf("Part 1: expected %d, got %d", part1Expected, part1Result)
	}

	part2Result := day2(strings.NewReader(TestInput), 2)
	if part2Result != part2Expected {
		t.Errorf("Part 2: expected %d, got %d", part2Expected, part2Result)
	}
}

func BenchmarkDay2(b *testing.B) {
	for b.Loop() {
		day2(strings.NewReader(TestInput), 1)
		day2(strings.NewReader(TestInput), 2)
	}
}

// 19128774598
