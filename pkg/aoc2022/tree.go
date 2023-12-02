package aoc2022

import (
	"bufio"
	"fmt"
	"strings"
)

func SeeTrees(b []byte) (int, int) {
	// trees will be identified by "row,column"

	// seeTree from outside the forest
	seeTree := make(map[string]struct{})
	// treeScenicScore sums the scenicScore of each tree
	treeScenicScore := make(map[string]int)
	var currentTreeRow int
	treeColumns := make(map[int][]rune)
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	var bestScenicScore int
	for scanner.Scan() {
		inputLine := scanner.Text()
		currentRow := []rune(inputLine) // use runes again, digits are 48-57
		var highFromLeft, highFromRight rune
		for currentTreeColumn, currentTreeHeight := range currentRow {
			currentTree := fmt.Sprintf("%d,%d", currentTreeRow, currentTreeColumn)
			// add to our column map for later processing
			treeColumns[currentTreeColumn] = append(treeColumns[currentTreeColumn], currentTreeHeight)

			// can we see tree from outside the forrest
			if currentTreeHeight > highFromLeft {
				seeTree[fmt.Sprintf("%d,%d", currentTreeRow, currentTreeColumn)] = struct{}{}
				highFromLeft = currentTreeHeight
			}

			treeFromRightColumn := len(currentRow) - 1 - currentTreeColumn
			treeFromRight := currentRow[treeFromRightColumn]
			if treeFromRight > highFromRight {
				seeTree[fmt.Sprintf("%d,%d", currentTreeRow, treeFromRightColumn)] = struct{}{}
				highFromRight = treeFromRight
			}

			// calculate row scenicScore per tree
			var scenicScore int
			treeScenicScore[currentTree] = 1

			// go right
			for i := currentTreeColumn + 1; i < len(currentRow); i++ {
				scenicScore++
				if currentRow[i] >= currentTreeHeight {
					break
				}
			}
			treeScenicScore[currentTree] = treeScenicScore[currentTree] * scenicScore
			scenicScore = 0

			// then go left
			for i := currentTreeColumn - 1; i >= 0; i-- {
				scenicScore++
				if currentRow[i] >= currentTreeHeight {
					break
				}
			}
			treeScenicScore[currentTree] = treeScenicScore[currentTree] * scenicScore

		}
		currentTreeRow++
	}

	// we have our tree columns mapped out
	// this is slow as heck, definitely could be optimized
	for currentTreeColumnIdx, col := range treeColumns {
		var highFromTop, highFromBottom rune
		for currentTreeRow, currentTreeHeight := range col {
			currentTree := fmt.Sprintf("%d,%d", currentTreeRow, currentTreeColumnIdx)

			if currentTreeHeight > highFromTop {
				seeTree[currentTree] = struct{}{}
				highFromTop = currentTreeHeight
			}
			treeRowFromBottom := len(col) - 1 - currentTreeRow
			treeHeightFromBottom := col[treeRowFromBottom]
			if treeHeightFromBottom > highFromBottom {
				seeTree[fmt.Sprintf("%d,%d", treeRowFromBottom, currentTreeColumnIdx)] = struct{}{}
				highFromBottom = treeHeightFromBottom
			}

			var scenicScore int
			// go down
			for i := currentTreeRow + 1; i < len(col); i++ {
				scenicScore++
				if col[i] >= currentTreeHeight {
					break
				}
			}
			treeScenicScore[currentTree] = treeScenicScore[currentTree] * scenicScore
			scenicScore = 0

			// then go up
			for i := currentTreeRow - 1; i >= 0; i-- {
				scenicScore++
				if col[i] >= currentTreeHeight {
					break
				}
			}
			treeScenicScore[currentTree] = treeScenicScore[currentTree] * scenicScore

			if treeScenicScore[currentTree] > bestScenicScore {
				bestScenicScore = treeScenicScore[currentTree]
			}

		}
	}

	return len(seeTree), bestScenicScore
}
