package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Day1() {
	r := bufio.NewReader(os.Stdin)
	i, err := SumCalibrationValues(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("----Solution----\n%v\n----------------", i)
}

// GetCalibrationValue extracts the first and last digit from a string,
// combines them into a two-digit number, and returns that number.
func GetCalibrationValue(s string) int {
	firstDigit := -1
	lastDigit := -1
	p := Part()

	switch p {
	case 1:
		for _, r := range s {
			if '0' <= r && r <= '9' {
				if firstDigit == -1 {
					firstDigit = int(r - '0')
				}
				lastDigit = int(r - '0')
			}
		}
		var calibrationValue int
		if firstDigit != -1 && lastDigit != -1 {
			calibrationValue = firstDigit*10 + lastDigit
		}
		return calibrationValue
	case 2:
		return PartTwoCalibration(s)
	default:
		fmt.Println("invalid part")

	}
	return -1
}

// SumCalibrationValues reads lines from an io.Reader, calculates the calibration value for each line,
// and returns the sum of all calibration values.
func SumCalibrationValues(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		// log := Log.With(zap.String("line", line))
		// log.Info("scanned line")
		sum += GetCalibrationValue(line)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

func PartTwoCalibration(s string) int {
	// the first and last digits may be represented by their string such as one=1, two=2, etc.
	// or they may be represented by their integer such as 1, 2, etc.
	numberMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	firstDigit := -1
	firstDigitIdx := -1
	lastDigit := -1
	lastDigitIdx := -1

	for numStr, num := range numberMap {
		if i := strings.Index(s, numStr); i != -1 &&
			(firstDigitIdx == -1 || i < firstDigitIdx) {
			firstDigit = num
			firstDigitIdx = i
		}
		if i := strings.LastIndex(s, numStr); i > lastDigitIdx {
			lastDigit = num
			lastDigitIdx = i
		}
	}
	var calibrationValue int
	if firstDigit != -1 && lastDigit != -1 {
		calibrationValue = firstDigit*10 + lastDigit
	}
	return calibrationValue
}
