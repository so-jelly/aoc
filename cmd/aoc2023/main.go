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
	case 3:
		Day3()
	case 4:
		Day4()
	case 5:
		Day5()
	case 6:
		Day6()
	case 7:
		Day7()
	default:
		fmt.Println("invalid day")
		return
	}
}
