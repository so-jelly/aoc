package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Map struct {
	ranges  []Range
	numbers []int
	name    string
	source  string
	dest    string
	// kind of a quasi-linked list
	nextName string
}

type Range struct {
	dstStart int
	srcStart int
	repeats  int
}

func Day5() {
	processDay5(bufio.NewReader(os.Stdin))
	fmt.Println("data processed")
	day5()
}

var maps = make(map[string]*Map)

func processDay5(r io.Reader) {
	scanner := bufio.NewScanner(r)
	current := &Map{
		ranges:   []Range{},
		name:     "seeds",
		source:   "",
		dest:     "seed",
		nextName: "",
	}
	maps["seeds"] = current

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			seeds := getSeeds(line)
			maps["seeds"].numbers = seeds
			continue
		}
		if strings.HasSuffix(line, "map:") {
			previousMap := current.name
			current = NewMap(line)
			maps[current.name] = current
			maps[previousMap].nextName = current.name
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		current.ranges = append(current.ranges, NewRange(line))
	}
}

func getSeeds(input string) []int {
	part := Part() // make sure we have this variable
	input = strings.Split(input, ":")[1]
	inputStrings := strings.Fields(input)
	ints := make([]int, 0, len(inputStrings))
	for _, in := range inputStrings {
		out, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		ints = append(ints, out)
	}
	if part == 1 {
		return ints
	}
	if len(ints)%2 != 0 {
		panic(fmt.Errorf("part 2 needs 2-tuples: %s %v", input, ints))
	}
	seeds := make([]int, 0)
	for i := 0; i < len(ints); i += 2 {
		seeds = append(seeds, makeRange(ints[i], ints[i+1])...)
	}
	// fmt.Println("found seeds:", seeds)
	return seeds
}

func makeRange(start, count int) []int {
	fmt.Println("makeRange:", start, count)
	numbers := make([]int, 0) // what's faster, create with capacity or append?
	if count < 0 {
		return numbers
	}
	for i := 0; i < count; i++ {
		numbers = append(numbers, start+i)
	}

	// fmt.Println("range:", numbers)
	return numbers
}

func NewRange(line string) Range {
	r := Range{}
	parts := strings.Fields(line)
	var err error
	if len(parts) != 3 {
		panic(fmt.Errorf("invalid line: %s", line))
	}
	r.dstStart, err = strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	r.srcStart, err = strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	r.repeats, err = strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	return r
}

func (r Range) InRange(n int) int {
	if n < r.srcStart || n > r.srcStart+r.repeats {
		return -1
	}
	return r.dstStart + n - r.srcStart
}

func NewMap(line string) *Map {
	r := regexp.MustCompile(`((\w+)-to-(\w+)) map:`)
	matches := r.FindStringSubmatch(line)
	if len(matches) != 4 {
		panic(fmt.Errorf("invalid map line: %s", line))
	}
	m := new(Map)
	m.name = matches[1]
	m.source = matches[2]
	m.dest = matches[3]
	return m
}

func (m *Map) AddRange(r Range) { m.ranges = append(m.ranges, r) }

func day5() {
	var lastMap, nextMap *Map
	currentMap := maps["seeds"]

	for {
		lastMap = currentMap
		if currentMap.nextName == "" {
			break
		}
		nextMap = maps[currentMap.nextName]
		numbers := make([]int, 0, len(currentMap.numbers))

		for _, n := range currentMap.numbers {
			numbers = append(numbers, getNewLocation(currentMap.ranges, n))
		}

		nextMap.numbers = numbers
		currentMap = nextMap
	}
	fmt.Println(getLowest(lastMap.numbers))
}

func getLowest(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // or return an error
	}

	lowest := numbers[0]
	for _, num := range numbers[1:] {
		if num < lowest {
			lowest = num
		}
	}
	return lowest
}

func getNewLocation(ranges []Range, n int) int {
	for _, r := range ranges {
		if newLocation := r.InRange(n); newLocation != -1 {
			return newLocation
		}
	}
	return n
}
