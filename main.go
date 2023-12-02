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
	default:
		fmt.Println("invalid day selected or not implemented")
		fmt.Printf("json encoding %s\n", jsonEscape(string(in)))
	}

}
