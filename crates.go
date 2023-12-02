package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type crateStack []string

// Place a crate on a stack
func (s *crateStack) Place(c crateStack) {
	*s = append(*s, c...) // Simply append the new value to the end of the stack
}

// Pick a crate from a stack
func (s *crateStack) Pick(n int) crateStack {
	index := len(*s) - n
	crates := (*s)[index:]
	*s = (*s)[:index]
	return crates
}

func Arrange(b []byte, version int) string {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	var doneCrates bool
	var reverseCrates int // 0 not ready, 1 reverse, -1 done reversing
	positions := make([]*crateStack, 1)

	// move 1 from 2 to 1
	moveRegex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	for scanner.Scan() {
		line := scanner.Text()
		linerunes := []rune(line)
		if len(linerunes) == 0 {
			continue
		}

		if !doneCrates {
			if string(linerunes[1]) == "1" {
				fmt.Println("done mapping crates before reversing")
				for _, pos := range positions {
					fmt.Println(pos)
				}
				fmt.Println("---")

				doneCrates = true
				reverseCrates = 1
				continue
			}
			for i, r := range linerunes {
				crate := string(r)
				if crate == " " {
					continue
				}
				if i == 1 {
					if positions[0] == nil {
						positions[0] = &crateStack{}
					}
					*positions[0] = append(*positions[0], crate)
				}
				if i > 4 && (i-1)%4 == 0 {
					position := ((i - 1) / 4) // 1=0,5=1,9=2
					for len(positions) <= position {
						positions = append(positions, &crateStack{})
					}
					*positions[position] = append(*positions[position], crate)
				}
			}
		}

		// reverse the crate stacks so we can pop and push from the end
		if reverseCrates == 1 {
			for i, cs := range positions {
				stack := *cs
				bottom := 0
				top := len(stack) - 1
				for bottom < top {
					stack[bottom], stack[top] = stack[top], stack[bottom]
					bottom++
					top--
				}
				*positions[i] = stack

			}
			reverseCrates = -1
			fmt.Println("reversed")
			for _, pos := range positions {
				fmt.Println(*pos)
			}
			fmt.Println("---")

		}

		if reverseCrates == -1 {
			moveQty, err := strconv.Atoi(moveRegex.ReplaceAllString(line, "$1"))
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(moveRegex.ReplaceAllString(line, "$2"))
			if err != nil {
				panic(err)
			}
			from--
			to, err := strconv.Atoi(moveRegex.ReplaceAllString(line, "$3"))
			if err != nil {
				panic(err)
			}
			to--

			if version == 9000 {
				for i := 0; i < moveQty; i++ {
					crate := positions[from].Pick(1)
					positions[to].Place(crate)
				}
			}

			if version == 9001 {
				positions[to].Place(positions[from].Pick(moveQty))
			}

		}
	}

	topCrates := ""
	fmt.Println("final config")
	for pos, stack := range positions {
		fmt.Println(pos, *stack)
		topCrates = topCrates + stack.Pick(1)[0]
	}
	fmt.Println("---")

	return topCrates
}
