package aoc2023

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{10, 5, 5},
		{15, 10, 5},
		{21, 14, 7},
		{36, 48, 12},
		{17, 23, 1},
	}

	for _, tt := range tests {
		got := gcd(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("gcd(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
func TestLCM(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{10, 5, 10},
		{15, 10, 30},
		{21, 14, 42},
		{36, 48, 144},
		{17, 23, 391},
		{2438, 1896, 2311224},
		{2311224, 17885, 41336241240},
	}

	for _, tt := range tests {
		got := lcm(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("lcm(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLCMM(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{10, 5}, 10},
		{[]int{15, 10}, 30},
		{[]int{21, 14}, 42},
		{[]int{36, 48}, 144},
		{[]int{17, 23}, 391},
		{[]int{2438, 1896}, 2311224},
		{[]int{2311224, 17885}, 41336241240},
		{[]int{2438, 1896, 17885}, 41336241240},
	}

	for _, tt := range tests {
		got := lcmm(tt.numbers)
		if got != tt.want {
			t.Errorf("lcmm(%v) = %d; want %d", tt.numbers, got, tt.want)
		}
	}
}
