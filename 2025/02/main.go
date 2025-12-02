package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

func onComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[:i], nil
	}
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func part1Valid(id string) bool {
	half := len(id) / 2
	return id[:half] != id[half:]
}

func part2Valid(id string) bool {
	n := len(id)
	for i := 1; i <= n/2; i++ { // only need to check half of the chunks
		if n%i != 0 {
			continue
		}
		if id[:n-i] == id[i:] { // 824824824 div 3 (not 2)
			return false
		}
	}
	return true
}

func day2(data io.Reader, part int) (sum int) {
	scanner := bufio.NewScanner(data)
	scanner.Split(onComma)

	for scanner.Scan() {
		valueRange := strings.TrimSpace(scanner.Text())
		if valueRange == "" {
			continue
		}
		idRange := strings.Split(valueRange, "-")
		startId, err := strconv.Atoi(idRange[0])
		if err != nil {
			panic(err)
		}
		endId, err := strconv.Atoi(idRange[1])
		if err != nil {
			panic(err)
		}

		id := startId
		for {
			if part == 1 {
				if !part1Valid(strconv.Itoa(id)) {
					sum = sum + id
				}
			} else {
				if !part2Valid(strconv.Itoa(id)) {
					sum = sum + id
				}
			}
			if id == endId {
				break
			}
			id++
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
		fmt.Println("part1", day2(part1reader, 1))
	}()
	go func() {
		defer wg.Done()
		defer pipeWriter.Close()
		fmt.Println("part2", day2(part2reader, 2))
	}()
	wg.Wait()

}
