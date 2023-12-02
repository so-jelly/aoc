package aoc2022

import "testing"

var crateTestData = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestArrange(t *testing.T) {
	stackTops9000 := Arrange([]byte(crateTestData), 9000)
	if stackTops9000 != "CMZ" {
		t.Errorf("Expected 'CMZ', got '%s'", stackTops9000)
	}

	stackTops9001 := Arrange([]byte(crateTestData), 9001)
	if stackTops9001 != "MCD" {
		t.Errorf("Expected 'MCD', got '%s'", stackTops9001)
	}

}
