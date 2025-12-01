package aoc2022

import (
	"bufio"
	"strconv"
	"strings"
)

func BridgeStep(b []byte, knots int) int {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))

	tailVisited := make(map[int]map[int]struct{})

	knotPositions := make(map[int][]int, knots)
	for i := 1; i <= knots; i++ {
		knotPositions[i] = []int{0, 0}
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		move, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		// xOrY is the move direction, position index slice
		// eg: headPos[xOrY] = n
		var xOrY int
		switch line[0] {
		case "R":
			xOrY = 0
		case "L":
			xOrY = 0
			move = -move
		case "U":
			xOrY = 1
		case "D":
			xOrY = 1
			move = -move
		}
		// move head
		moveHeadRemaining, moveHead := AbsMove(move)
		for i := 0; i < moveHeadRemaining; i++ {
			knotPositions[1][xOrY] = knotPositions[1][xOrY] + moveHead
			for knot := 2; knot <= knots; knot++ {
				knotMove := CalculateMove(knotPositions[knot], knotPositions[knot-1])
				knotPositions[knot][0] = knotPositions[knot][0] + knotMove[0]
				knotPositions[knot][1] = knotPositions[knot][1] + knotMove[1]
			}
			if tailVisited[knotPositions[knots][0]] == nil {
				tailVisited[knotPositions[knots][0]] = make(map[int]struct{})
			}
			tailVisited[knotPositions[knots][0]][knotPositions[knots][1]] = struct{}{}
		}
	}
	var tvCount int
	for _, v := range tailVisited {
		tvCount += len(v)
	}
	return tvCount
}

// AbsMove returns the absolute value and the move to get closer
func AbsMove(x int) (int, int) {
	if x > 0 {
		return x, 1
	}
	if x < 0 {
		return -x, -1
	}
	return 0, 0
}

// CalculateMove returns the absolute
func CalculateMove(tail, head []int) []int {
	tailXtoHead := head[0] - tail[0]
	tailYtoHead := head[1] - tail[1]
	xToHead, moveX := AbsMove(tailXtoHead)
	yToHead, moveY := AbsMove(tailYtoHead)
	if (xToHead > 1 && yToHead > 0) ||
		(xToHead > 0 && yToHead > 1) {
		return []int{moveX, moveY}
	}
	if xToHead > 1 {
		return []int{moveX, 0}
	}
	if yToHead > 1 {
		return []int{0, moveY}
	}
	return []int{0, 0}
}
