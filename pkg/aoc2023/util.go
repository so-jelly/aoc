package aoc2023

import (
	"fmt"
	"os"
	"strconv"
)

func Day() int  { return GetEnvInt("DAY") }
func Part() int { return GetEnvInt("PART") }

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
