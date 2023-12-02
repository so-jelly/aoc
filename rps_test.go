package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

const (
	RPStestdata    = "testdata/testrps.json"
	RPSbadTestData = "testdata/testrps_bad.json"
)

type TestRPSData []TestRPSDatum

type TestRPSDatum struct {
	Input        string `json:"input"`
	Score        int    `json:"score"`
	OptimalScore int    `json:"optimalScore"`
}

func getRPStest(fileName string) *TestRPSData {
	in, _ := os.ReadFile(RPStestdata)
	testData := &TestRPSData{}
	json.Unmarshal(in, &testData)
	return testData
}

func TestScoreRPStournament(t *testing.T) {
	testData := getRPStest(RPStestdata)
	for _, testCase := range *testData {
		input := []byte(testCase.Input)
		haveScore, haveOptimalScore := ScoreRPStournament(input)
		wantScore, wantOptimalScore := testCase.Score, testCase.OptimalScore
		fmt.Println(haveScore, haveOptimalScore)
		if haveScore != wantScore ||
			haveOptimalScore != wantOptimalScore {
			t.Errorf("have score %d, optimalscore: %d\nwant score %d, optimalscore:%d\n", haveScore, haveOptimalScore, wantScore, wantOptimalScore)
		}
	}

}
