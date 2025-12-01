package aoc2023

import (
	"fmt"
	"testing"
)

func TestRange_InRange(t *testing.T) {
	tests := []struct {
		r        Range
		n        int
		expected int
	}{
		{
			r:        Range{srcStart: 98, repeats: 99, dstStart: 50},
			n:        90,
			expected: -1,
		},
		{
			r:        Range{srcStart: 98, repeats: 99, dstStart: 50},
			n:        98,
			expected: 50,
		},
		{
			r:        Range{srcStart: 98, repeats: 99, dstStart: 50},
			n:        99,
			expected: 51,
		},
		// Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
		{Range{50, 98, 2}, 79, -1},
		{Range{52, 50, 48}, 79, 81},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := test.r.InRange(test.n)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}

func TestMakeRange(t *testing.T) {
	tests := []struct {
		start    int
		count    int
		expected []int
	}{
		{
			start:    1,
			count:    5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			start:    10,
			count:    15,
			expected: []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		},
		{
			start:    -3,
			count:    0,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("start=%d,count=%d", test.start, test.count), func(t *testing.T) {
			result := makeRange(test.start, test.count)
			if len(result) != test.count {
				t.Errorf("Expected length %d, but got %d", test.count, len(result))
			}
			for i := 0; i < len(result); i++ {
				if result[i] != test.expected[i] {
					t.Errorf("Expected %d, but got %d", test.expected[i], result[i])
				}
			}
		})
	}
}
