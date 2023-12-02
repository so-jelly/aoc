package main

import (
	"bufio"
	"strconv"
	"strings"
)

func ParseElfSections(s string) (low, high int) {
	// 1-2
	e := strings.Split(s, "-")
	low, err := strconv.Atoi(e[0])
	if err != nil {
		panic(err)
	}
	high, err = strconv.Atoi(e[1])
	if err != nil {
		panic(err)
	}
	return low, high
}

func FullOverlaps(b []byte) int {
	var overlaps int
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ",")
		elfOneLow, elfOneHigh := ParseElfSections(a[0])
		elfTwoLow, elfTwoHigh := ParseElfSections(a[1])
		if elfOneLow <= elfTwoLow && elfOneHigh >= elfTwoHigh {
			// fmt.Println("1", a, elfOne, elfTwo)
			overlaps++
			continue // in case of a double match
		}
		if elfTwoLow <= elfOneLow && elfTwoHigh >= elfOneHigh {
			// fmt.Println("2", a, elfOne, elfTwo)
			overlaps++
		}
	}
	return overlaps
}

func PartialOverlaps(b []byte) int {
	var overlaps int
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ",")
		elfOneLow, elfOneHigh := ParseElfSections(a[0])
		elfTwoLow, elfTwoHigh := ParseElfSections(a[1])
		// 1-3,2-4
		if elfOneLow <= elfTwoLow &&
			elfOneHigh >= elfTwoLow {
			overlaps++
			continue
		}
		if elfTwoLow <= elfOneLow &&
			elfTwoHigh >= elfOneLow {
			overlaps++
		}
	}
	return overlaps
}
