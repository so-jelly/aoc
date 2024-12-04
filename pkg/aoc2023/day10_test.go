package aoc2023

import (
	"reflect"
	"testing"
)

func TestInboundConnections(t *testing.T) {
	grid := [][]string{
		{SW, SE, SW},
		{SE, S, NS},
		{NE, EW, NW},
	}

	tests := []struct {
		grid     [][]string
		position []int
		expected [][]int
	}{
		{
			position: []int{0, 0},
			expected: nil,
		},
		{
			position: []int{1, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			position: []int{2, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			position: []int{1, 0},
			expected: [][]int{{2, 0}, {1, 1}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := InboundConnections(grid, test.position)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
