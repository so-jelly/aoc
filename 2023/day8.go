package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func init() {
	DayFunc[8] = Day8
}

func Day8(part int, r io.Reader) {
	stime := time.Now()
	fmt.Println("Day 8")
	d8processData(r)
	// fmt.Println(instructions, d8nodes)
	if part == 1 {
		d8travel()
	} else {
		d8p2travel()
	}
	fmt.Printf("%dns\n", time.Since(stime).Nanoseconds())

}

var instructions []int // 0 for left, 1 for right
var d8nodes = make(map[string][]string)

func d8processData(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if instructions == nil {
			for _, r := range line {
				i := 0
				if string(r) == "R" {
					i = 1
				}
				instructions = append(instructions, i)
			}
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		nodeStr := strings.Split(line, " = ")
		node := nodeStr[0]
		destNodes := make([]string, 2)
		destNodesFields := strings.Fields(nodeStr[1])
		destNodes[0] = strings.Trim(destNodesFields[0], "(, ")
		destNodes[1] = strings.Trim(destNodesFields[1], ") ")
		d8nodes[node] = destNodes
	}
}

func d8travel() {
	node := "AAA"
	var instructionIndex, steps int
	for {
		steps++
		node = d8nodes[node][instructions[instructionIndex]]
		// fmt.Println("move to ", node)
		if node == "ZZZ" {
			fmt.Println(steps)
			return
		}
		instructionIndex++
		if instructionIndex > len(instructions)-1 {
			instructionIndex = 0
		}
	}
}

func d8p2travel() {
	var p2nodes []string
	p2endNodes := make(map[string]struct{})
	for node := range d8nodes {
		if strings.HasSuffix(node, "A") {
			p2nodes = append(p2nodes, node)
		}
		if strings.HasSuffix(node, "Z") {
			p2endNodes[node] = struct{}{}
		}
	}

	leastSteps := make([]int, len(p2nodes))
	var wg sync.WaitGroup
	for nodeIdx, n := range p2nodes {
		node := n
		wg.Add(1)
		go func(nodeIdx int, node string) {
			defer wg.Done()
			var steps, instructionIndex int
			for {
				steps++
				node = d8nodes[node][instructions[instructionIndex]]
				if _, ok := p2endNodes[node]; ok {
					fmt.Println("found end node", nodeIdx)
					leastSteps[nodeIdx] = steps
					break
				}
				instructionIndex++
				if instructionIndex > len(instructions)-1 {
					instructionIndex = 0
				}
			}
		}(nodeIdx, node)
	}
	wg.Wait()
	fmt.Println(leastSteps)
	fmt.Println(lcmm(leastSteps))
}
