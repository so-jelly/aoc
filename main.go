package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var dayFlag = flag.Int("d", 0, "adevent of code day")

func main() {

	flag.Parse()

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// to gather test data
	// fmt.Printf("json encoding: %s", jsonEscape(string(in)))

	fmt.Printf("evaluating for day %d!\n", *dayFlag)
	switch day := dayFlag; *day {
	case 1:
		e := GetElfCalories(in)
		fmt.Printf("max: %d\n", e.MaxCalories())
		fmt.Printf("sum top three: %d\n", e.SumTopThree())
	case 2:
		score, optimalScore := ScoreRPStournament(in)
		fmt.Printf("your score is %d\nyour optimal score is: %d\n", score, optimalScore)
	case 3:
		fmt.Printf("bag prio: %d\n", BagPriorities(in))
		fmt.Printf("group prio: %d\n", GroupPriorities(in))
	case 4:
		fmt.Printf("section full overlaps: %d\n", FullOverlaps(in))
		fmt.Printf("section partial overlaps: %d\n", PartialOverlaps(in))
	case 5:
		fmt.Printf("top crates on each stack after moving %s version 9000\n", Arrange(in, 9000))
		fmt.Printf("top crates on each stack after moving %s version 9001\n", Arrange(in, 9001))
	case 6:
		fmt.Printf("packet start found at %d\n", PacketStart(in, 4))
		fmt.Printf("message start found at %d\n", PacketStart(in, 14))
	case 7:
		smallDirs, delDirSize := DirSize(in, 100000)
		fmt.Printf("small dirs %v \ndelete dir size %v\n", smallDirs, delDirSize)
	case 8:
		seeTrees, scenicScore := SeeTrees(in)
		fmt.Printf("i see %d trees\nthe best scenic score is %d\n", seeTrees, scenicScore)
	default:
		fmt.Println("invalid day selected or not implemented")
		fmt.Printf("json encoding %s\n", jsonEscape(string(in)))
	}

}
