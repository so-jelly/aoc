package aoc2022

import "testing"

var sectionTestData = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestFullOverlaps(t *testing.T) {
	overlaps := FullOverlaps([]byte(sectionTestData))
	if overlaps != 2 {
		t.Errorf("Expected %d overlaps, got %d", 2, overlaps)
	}
}

func TestPartialOverlaps(t *testing.T) {
	overlaps := PartialOverlaps([]byte(sectionTestData))
	if overlaps != 4 {
		t.Errorf("Expected %d overlaps, got %d", 2, overlaps)
	}
}
