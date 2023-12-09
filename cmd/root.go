package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/so-jelly/aoc/pkg/aoc"
	aoc2023 "github.com/so-jelly/aoc/pkg/aoc2023"
	"github.com/spf13/cobra"
)

var year, day, part int
var getData, test bool

func init() {
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "aoc year")
	err := rootCmd.MarkFlagRequired("year")
	if err != nil {
		panic(err)
	}
	rootCmd.Flags().IntVarP(&day, "day", "d", 0, "aoc day")
	err = rootCmd.MarkFlagRequired("day")
	if err != nil {
		panic(err)
	}
	rootCmd.Flags().IntVarP(&part, "part", "p", 0, "aoc part")
	err = rootCmd.MarkFlagRequired("part")
	if err != nil {
		panic(err)
	}
	rootCmd.Flags().BoolVarP(&getData, "get", "g", false, "get input data from aoc website")
	rootCmd.Flags().BoolVarP(&test, "test", "t", false, "run against example input")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		if getData {
			fmt.Printf("advent of code. year: %d, day: %d, part: %d\n", year, day, part)
			err := aoc.Get(year, day, part)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("input data saved to data folder")
			return
		}

		inputFile := fmt.Sprintf("data/%d/%d/input.txt", year, day)
		if test {
			inputFile = fmt.Sprintf("data/%d/%d/example_input.txt", year, day)
		}

		r, err := os.Open(inputFile)
		if err != nil {
			log.Printf("failed to open input file (may need to run `task get`): %v", err)
			return
		}

		switch year {
		case 2023:
			if _, ok := aoc2023.DayFunc[day]; !ok {
				fmt.Println("invalid day")
				return
			}
			aoc2023.DayFunc[day](part, r)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
