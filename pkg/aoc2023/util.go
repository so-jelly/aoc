package aoc2023

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var DayFunc = make(map[int]func(int, io.Reader))

func GetEnvInt(s string) int {
	e := os.Getenv(s)
	if e == "" {
		fmt.Printf("environment variable %s is empty, defaulting to 1\n", s)
		return 1
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		fmt.Printf("error converting environment variable %s to int\n", s)
	}
	return i
}

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		var temp = b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a, b int) int {
	return (a * b / gcd(a, b))
}

func lcmm(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}
	if len(numbers) == 2 {
		return lcm(numbers[0], numbers[1])
	}
	return lcm(numbers[0], lcmm(numbers[1:]))
}
