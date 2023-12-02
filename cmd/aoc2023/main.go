package main

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/so-jelly/aoc/pkg/aoc2023"
)

var day int

func main() {
	dayEnv := os.Getenv("DAY")
	day, _ = strconv.Atoi(dayEnv) // convert string to int, ignoring error for simplicity

	switch day {
	case 1:
		Day1()
		return
	default:
		fmt.Println("invalid day")
		return
	}
}
