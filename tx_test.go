package main

import "testing"

func TestPacketStart(t *testing.T) {

	testData := map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for data, result := range testData {
		if r := PacketStart([]byte(data), 4); r != result {
			t.Errorf("Expected %d, got %d", result, r)
		}
	}

	testMessageData := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	for data, result := range testMessageData {
		if r := PacketStart([]byte(data), 14); r != result {
			t.Errorf("Expected %d, got %d", result, r)
		}
	}

}
