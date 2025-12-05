package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

const tp byte = '@'

// const notTp byte = '.'

func copyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func processPass(s io.Reader, w io.Writer) int {
	scanner := bufio.NewScanner(s)
	var prev, curr, next []byte
	removed := 0

	process := func(p, c, n []byte) {
		for i := range c {
			if c[i] != tp {
				if w != nil {
					w.Write([]byte{c[i]})
				}
				continue
			}
			var adjacent int
			// Check 8 neighbors
			if p != nil {
				if i > 0 && p[i-1] == tp {
					adjacent++
				}
				if p[i] == tp {
					adjacent++
				}
				if i < len(p)-1 && p[i+1] == tp {
					adjacent++
				}
			}
			if i > 0 && c[i-1] == tp {
				adjacent++
			}
			if i < len(c)-1 && c[i+1] == tp {
				adjacent++
			}
			if n != nil {
				if i > 0 && n[i-1] == tp {
					adjacent++
				}
				if n[i] == tp {
					adjacent++
				}
				if i < len(n)-1 && n[i+1] == tp {
					adjacent++
				}
			}

			if adjacent < 4 {
				removed++
				if w != nil {
					w.Write([]byte{'.'})
				}
			} else {
				if w != nil {
					w.Write([]byte{tp})
				}
			}
		}
		if w != nil {
			w.Write([]byte{'\n'})
		}
	}

	if scanner.Scan() {
		curr = copyBytes(scanner.Bytes())
	}

	for scanner.Scan() {
		next = copyBytes(scanner.Bytes())
		process(prev, curr, next)
		prev = curr
		curr = next
	}
	if curr != nil {
		process(prev, curr, nil)
	}
	return removed
}

func part1(s io.Reader) (forkliftPositions int) {
	return processPass(s, nil)
}

func part2(s io.Reader) int {
	totalRemoved := 0
	currentInput := s

	for {
		output := new(bytes.Buffer)
		removedInPass := processPass(currentInput, output)

		if removedInPass == 0 {
			break
		}
		totalRemoved += removedInPass
		currentInput = output
	}
	return totalRemoved
}

func day4(data io.Reader, part int) (sum int) {
	if part == 1 {
		sum += part1(data)
	} else {
		sum += part2(data)
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
		fmt.Println("part1", day4(part1reader, 1))
	}()
	go func() {
		defer wg.Done()
		defer pipeWriter.Close()
		fmt.Println("part2", day4(part2reader, 2))
	}()
	wg.Wait()

}
