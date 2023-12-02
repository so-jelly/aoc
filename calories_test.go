package main

import (
	"encoding/json"
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

const (
	testCaloriesData = "testdata/testcalories.json"
)

type TestCaloriesData []TestCaloriesDatum

type TestCaloriesDatum struct {
	Input       string `json:"input"`
	Calories    []int  `json:"calories"`
	Max         int    `json:"max"`
	SumTopThree int    `json:"sumTopThree"`
}

func getTestCalories(fileName string) *TestCaloriesData {
	in, _ := os.ReadFile(testCaloriesData)
	testData := &TestCaloriesData{}
	json.Unmarshal(in, &testData)
	return testData
}

func TestGetCalories(t *testing.T) {
	testData := getTestCalories(testCaloriesData)
	for _, testCase := range *testData {
		input := []byte(testCase.Input)
		have := GetElfCalories(input)
		want := ElfCalories(testCase.Calories)
		if slices.Compare(want, *have) != 0 {
			t.Errorf("have %d, want %d", have, want)
		}
	}
}

func TestMaxCalories(t *testing.T) {
	testData := getTestCalories(testCaloriesData)
	for _, testCase := range *testData {
		input := []byte(testCase.Input)
		e := GetElfCalories(input)
		have := e.MaxCalories()
		want := testCase.Max
		if have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}

}

func TestSumCalories(t *testing.T) {
	testData := getTestCalories(testCaloriesData)
	for _, testCase := range *testData {
		input := []byte(testCase.Input)
		e := GetElfCalories(input)
		have := e.SumTopThree()
		want := testCase.SumTopThree
		if have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
}
