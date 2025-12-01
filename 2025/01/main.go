package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func day1(data io.Reader, part int) (password int) {
	var position = 50
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var lr string
		var num int
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%1s%d", &lr, &num)
		if err != nil {
			fmt.Println("Error scanning line:", line)
			panic(err)
		}
		if part == 2 {
			password += num / 100
		}
		remainder := num % 100
		if lr == "L" {
			wasZero := position == 0
			position = position - remainder
			if position < 0 {
				position = 100 + position
				if part == 2 && position != 0 && !wasZero {
					password++
				}
			}
		} else {
			position = position + remainder
			if position >= 100 {
				position = position - 100
				if part == 2 && position != 0 {
					password++
				}
			}
		}
		if position == 0 {
			password++
		}
	}
	return password
}

func main() {
	part1reader, pipeWriter := io.Pipe()
	part2reader := io.TeeReader(os.Stdin, pipeWriter)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("part1", day1(part1reader, 1))
	}()
	go func() {
		defer wg.Done()
		defer pipeWriter.Close()
		fmt.Println("part2", day1(part2reader, 2))
	}()
	wg.Wait()

}
