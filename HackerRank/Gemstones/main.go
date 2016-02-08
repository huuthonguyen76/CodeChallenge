// Challenge Url: https://www.hackerrank.com/challenges/gem-stones

package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalCases int
	var cases = make([]string, 0)
	var currentScanString string
	var result int

	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentScanString)
		cases = append(cases, currentScanString)
	}
	result = countGemStones(cases)

	fmt.Println(result)
}

func countGemStones(cases []string) int {
	var currentCase []string
	var listGemStones = make(map[string]int)
	var result = 0

	for i := 0; i < len(cases); i++ {
		currentCase = strings.Split(cases[i], "")
		currentTrack := make(map[string]bool)

		for j := 0; j < len(currentCase); j++ {
			if _, isOk := listGemStones[currentCase[j]]; !isOk {
				listGemStones[currentCase[j]] = 0
			}

			if _, isOk := currentTrack[currentCase[j]]; !isOk {
				listGemStones[currentCase[j]] += 1
				currentTrack[currentCase[j]] = true
			}
		}
	}

	for _, val := range listGemStones {
		if val == len(cases) {
			result++
		}
	}
	return result
}
