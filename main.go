package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("Need input")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	var elfCalories []int
	var currentElfCalories int

	for scanner.Scan() {
		if scanner.Text() == "" {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("input should be a list of integers")
			panic(err)
		}
		currentElfCalories += calories
	}
	// add the last elf
	elfCalories = append(elfCalories, currentElfCalories)

	var maxCalories int
	for _, calories := range elfCalories {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Println(maxCalories)
}
