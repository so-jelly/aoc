package aoc2022

import "encoding/json"

func JsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(b)
}
