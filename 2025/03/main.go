package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

const TestInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func part1(s []byte) int {
	firstDigit, firstDigitIndex := s[0], 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] > firstDigit {
			firstDigit = s[i]
			firstDigitIndex = i
		}
	}
	secondDigit := s[firstDigitIndex+1]
	for i := firstDigitIndex + 2; i < len(s); i++ {
		if s[i] > secondDigit {
			secondDigit = s[i]
		}
	}
	j, err := strconv.Atoi(string([]byte{firstDigit, secondDigit}))
	if err != nil {
		panic(err)
	}
	return j
}

func part2(availableBatteries []byte) int {
	batteriesNeeded := 12
	batteries := make([]byte, 0, batteriesNeeded)
	for availableBatteryIndex := range availableBatteries {
		battery := availableBatteries[availableBatteryIndex]
		for len(batteries) > 0 && batteries[len(batteries)-1] < battery && len(batteries)-1+(len(availableBatteries)-availableBatteryIndex) >= batteriesNeeded {
			batteries = batteries[:len(batteries)-1]
		}
		if len(batteries) < batteriesNeeded {
			batteries = append(batteries, battery)
		}
	}

	k, err := strconv.Atoi(string(batteries))
	if err != nil {
		panic(err)
	}
	return k
}

func day3(data io.Reader, part int) (sum int) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		// just used bytes, already a comparable slice of values
		bytes := scanner.Bytes()
		if part == 1 {
			sum += part1(bytes)
		} else {
			sum += part2(bytes)
		}
	}

	return sum
}

func main() {
	part1reader, pipeWriter := io.Pipe()
	part2reader := io.TeeReader(os.Stdin, pipeWriter)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("part1", day3(part1reader, 1))
	}()
	go func() {
		defer wg.Done()
		defer pipeWriter.Close()
		fmt.Println("part2", day3(part2reader, 2))
	}()
	wg.Wait()

}
