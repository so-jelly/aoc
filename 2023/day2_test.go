package aoc2023

import (
	"reflect"
	"testing"
)

func TestParseGame(t *testing.T) {
	s := "Game 1234: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	expectedGameId := 1234
	expectedPeek := []map[string]int{
		{"green": 1, "red": 3, "blue": 6},
		{"green": 3, "red": 6},
		{"green": 3, "blue": 15, "red": 14},
	}

	game := ParseGame(s)

	if game.Id != expectedGameId {
		t.Errorf("Expected game ID %d, but got %d", expectedGameId, game.Id)
	}

	if !reflect.DeepEqual(game.Peeks, expectedPeek) {
		t.Errorf("Expected peek %v, but got %v", expectedPeek, game.Peeks)
	}
}
