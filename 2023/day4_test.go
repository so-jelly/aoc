package aoc2023

// import (
// 	"fmt"
// 	"slices"
// 	"strings"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// )

// // func TestGetPoints(t *testing.T) {
// // 	tests := []struct {
// // 		input    int
// // 		expected int
// // 	}{
// // 		{0, 0},
// // 		{1, 1},
// // 		{2, 2},
// // 		{3, 4},
// // 		{4, 8},
// // 		{5, 16},
// // 		{14, 8192},
// // 		// Add more test cases as needed
// // 	}

// // 	for _, test := range tests {
// // 		output := getPoints(test.input)
// // 		if output != test.expected {
// // 			t.Errorf("getPoints(%d) = %d, want %d", test.input, output, test.expected)
// // 		}
// // 	}
// // }

// func Test_Day4(t *testing.T) {

// 	processData(strings.NewReader(input))
// 	var points int
// 	for _, c := range cards {
// 		c.FindWinningNumbers()
// 		c.CalculatePoints()
// 		points += c.points
// 		// .Printf("Card %+v\n", c)
// 	}
// 	// t.Errorf("%v", points)

// }

// func TestStringToIntSlice(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected []int
// 		err      error
// 	}{
// 		{input: " 33 13 28 76 16 91 52 41 38 64 ",
// 			expected: []int{33, 13, 28, 76, 16, 91, 52, 41, 38, 64},
// 			err:      nil,
// 		},
// 		{
// 			input:    "1 2 3 4 5",
// 			expected: []int{1, 2, 3, 4, 5},
// 			err:      nil,
// 		},
// 		{
// 			input:    "10 20 30",
// 			expected: []int{10, 20, 30},
// 			err:      nil,
// 		},
// 		{
// 			input:    "100",
// 			expected: []int{100},
// 			err:      nil,
// 		},
// 		{
// 			input:    "  5  10  15  ",
// 			expected: []int{5, 10, 15},
// 			err:      nil,
// 		},
// 		// {
// 		// 	input:    "1 a 3",
// 		// 	expected: nil,
// 		// 	err:      fmt.Errorf("invalid number: 'a' '1 a 3'"),
// 		// },
// 	}

// 	for _, test := range tests {
// 		result, err := stringToIntSlice(test.input)
// 		if !cmp.Equal(result, test.expected) || err != nil {
// 			t.Errorf("Input: %s\nExpected: %v, %v\nGot: %v, %v", test.input, test.expected, test.err, result, err)
// 		}
// 	}
// }
