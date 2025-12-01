package aoc2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	diskSize          = 70000000
	updateSpaceNeeded = 30000000
)

func DirSize(b []byte, maxSize int) (int, int) {
	pathUsage := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	currentPath := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		switch listingItem := strings.Split(line, " "); listingItem[0] {
		case "$":
			if listingItem[1] == "cd" {
				switch to := listingItem[2]; to {
				case "/":
					currentPath = []string{}
				case "..":
					currentPath = currentPath[:len(currentPath)-1]
				default:
					currentPath = append(currentPath, to)
				}
			}
		case "dir": // ignore dir listing
		default: // should be a size
			size, err := strconv.Atoi(listingItem[0])
			if err != nil {
				fmt.Println(line, size)
				panic(err)
			}

			for i := 0; i <= len(currentPath); i++ {
				pathString := strings.Join(currentPath[:i], "/")
				pathUsage[pathString] = pathUsage[pathString] + size
			}
		}
	}

	freeSpace := diskSize - pathUsage[""]
	needSpace := updateSpaceNeeded - freeSpace

	var sumSmallDirs, smallestDirThatWillFreeEnoughSpace int

	for _, v := range pathUsage {
		if v >= needSpace &&
			(smallestDirThatWillFreeEnoughSpace == 0 ||
				v < smallestDirThatWillFreeEnoughSpace) {
			smallestDirThatWillFreeEnoughSpace = v
		}

		if v <= maxSize {
			sumSmallDirs = sumSmallDirs + v
		}
	}

	return sumSmallDirs, smallestDirThatWillFreeEnoughSpace
}
