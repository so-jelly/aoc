package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func day5(data io.Reader) (totalFresh, freshIdCount int) {
	freshRanges := make([][2]int, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)
		freshRanges = append(freshRanges, [2]int{start, end})
	}

	slices.SortFunc(freshRanges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	writeIdx := 0
	for i := 1; i < len(freshRanges); i++ {
		curr := freshRanges[i]
		last := &freshRanges[writeIdx]
		if curr[0] <= last[1] { // extend
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
			// fully contained, do nothing
		} else {
			// No overlap: move to next slot and copy
			writeIdx++
			freshRanges[writeIdx] = curr
		}
	}
	// removed any extra ranges
	freshRanges = freshRanges[:writeIdx+1]

	var wg sync.WaitGroup

	wg.Go(func() {
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			for _, r := range freshRanges {
				if id >= r[0] && id <= r[1] {
					totalFresh++
					break
				}
			}
		}
	})

	wg.Go(func() {
		for _, r := range freshRanges {
			freshIdCount += r[1] - r[0] + 1
		}
	})

	wg.Wait()
	return totalFresh, freshIdCount
}

func main() {
	fmt.Println(day5(os.Stdin))
}
