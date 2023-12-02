package main

import (
	"fmt"

	. "github.com/so-jelly/aoc/pkg/aoc2023"
)

func main() {
	switch GetEnvInt("DAY") {
	case 1:
		Day1()
		return
	case 2:
		Day2()
		return
	default:
		fmt.Println("invalid day")
		return
	}
}