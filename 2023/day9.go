package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	DayFunc[9] = Day9
}

func Day9(part int, r io.Reader) {
	fmt.Println("Day 9")
	d9data(r)
	// fmt.Println(data)
	var next, previous int
	for _, d := range data {
		var nums [][]int
		nums = append(nums, d)
		nums = recursiveDifferences(nums)
		// fmt.Println("diff:", nums)
		nums = recursiveExtrapolate(nums)
		// fmt.Println("extrapolated", nums)
		next += nums[0][len(nums[0])-1]
		previous += nums[0][0]
	}
	fmt.Println("next (part1):", next)
	fmt.Println("previous (part2):", previous)
}

var data [][]int

func d9data(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		nums := make([]int, len(fields))
		for i, f := range fields {
			var err error
			nums[i], err = strconv.Atoi(f)
			if err != nil {
				panic(err)
			}
		}
		data = append(data, nums)
	}
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func getDifferences(nums []int) []int {
	var diffs []int
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	return diffs
}

func recursiveDifferences(nums [][]int) [][]int {
	differences := nums[len(nums)-1]
	if allZero(differences) {
		return nums
	}
	nums = append(nums, getDifferences(differences))
	return recursiveDifferences(nums)
}

func extrapolate(nums [][]int) [][]int {
	differences := nums[len(nums)-1]
	nums = nums[0 : len(nums)-1]
	extrapolate := nums[len(nums)-1]
	next := extrapolate[len(extrapolate)-1] + differences[len(differences)-1]
	extrapolate = append(extrapolate, next)
	previous := extrapolate[0] - differences[0]
	extrapolate = append([]int{previous}, extrapolate...)
	nums[len(nums)-1] = extrapolate
	return nums
}

func recursiveExtrapolate(nums [][]int) [][]int {
	if len(nums) == 1 {
		return nums
	}
	return recursiveExtrapolate(extrapolate(nums))
}
