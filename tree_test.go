package main

import "testing"

var treeTestData = `30373
25512
65332
33549
35390`

func TestSeeTrees(t *testing.T) {
	seeTrees, scenicScore := SeeTrees([]byte(treeTestData))
	if seeTrees != 21 {
		t.Errorf("I think I see %v, but I should see %v", seeTrees, 21)
	}
	if scenicScore != 8 {
		t.Errorf("Calculated scenic score %d, but I should have 8\n", scenicScore)
	}
}

var benchSeeTrees, benchScenicScore int

func BenchmarkSeeTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchSeeTrees, benchScenicScore = SeeTrees([]byte(treeTestData))
	}
}
