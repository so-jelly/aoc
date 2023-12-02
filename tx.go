package main

import (
	"bufio"
	"strings"
)

func PacketStart(b []byte, n int) int {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		line := scanner.Text()
		linerunes := []rune(line)

		for i := n - 1; i < len(linerunes)+1; i++ {
			unique := map[rune]bool{linerunes[i]: true}
			for j := 1; j <= n-1; j++ {
				k := linerunes[i-j]
				if _, val := unique[k]; val {
					break
				}
				unique[k] = true
				if j == n-1 {
					return i + 1
				}
			}

		}

	}
	return 0
}
