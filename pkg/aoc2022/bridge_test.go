package aoc2022

import (
	"slices"
	"testing"
)

func TestBridgeStep(t *testing.T) {

	testBridgeData := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	if steps := BridgeStep([]byte(testBridgeData), 2); steps != 13 {
		t.Errorf("BridgeStep should be 13, but I got %v", steps)
	}

	if steps := BridgeStep([]byte(testBridgeData), 10); steps != 1 {
		t.Errorf("BridgeStep should be 1, but I got %v", steps)
	}

	btest := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	if steps := BridgeStep([]byte(btest), 10); steps != 36 {
		t.Errorf("BridgeStep should be 36, but I got %v", steps)
	}

}

func TestAbsMove(t *testing.T) {
	d := [][]int{
		{0, 0, 0}, {1, 1, 1}, {-1, 1, -1},
		{2, 2, 1}, {-2, 2, -1},
	}
	for _, testD := range d {
		a, m := AbsMove(testD[0])
		if a != testD[1] || m != testD[2] {
			t.Errorf("abs test failed")
		}
	}
}

func TestCalculateMove(t *testing.T) {
	testData := [][][]int{
		{{0, 0}, {0, 0}, {0, 0}},     // no move
		{{0, 0}, {4, 0}, {1, 0}},     // move right
		{{3, 6}, {3, -11}, {0, -1}},  // move left
		{{0, 0}, {0, -4}, {0, -1}},   // move down
		{{1, 2}, {1, -4}, {0, -1}},   // move up
		{{12, 16}, {1, 1}, {-1, -1}}, // move diagonal
		{{2, 0}, {4, 0}, {1, 0}},     // move right
		{{4, 3}, {2, 4}, {-1, 1}},
	}
	for _, td := range testData {
		xy := CalculateMove(td[0], td[1])
		if slices.Compare(xy, td[2]) != 0 {
			t.Errorf("error calculating move %v %v: should be %v, got %v", td[0], td[1], td[2], xy)
		}
	}
}
