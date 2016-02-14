// Challenge Url: https://www.hackerrank.com/challenges/anagram

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var cases []string
	var currentCase string
	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentCase)
		cases = append(cases, currentCase)
	}

	for i := 0; i < len(cases); i++ {
		if len(cases[i])%2 != 0 {
			fmt.Println("-1")
		} else {
			firstString := cases[i][0 : len(cases[i])/2]
			secondString := cases[i][len(cases[i])/2 : len(cases[i])]

			fmt.Println(countDeletion(firstString, secondString))
		}
	}
}

func countDeletion(firstString string, secondString string) int {
	var listTrackFirst = make(map[uint8]int)
	var listTrackSecond = make(map[uint8]int)

	var totalEqual = 0

	for i := 0; i < len(firstString); i++ {
		if _, isOk := listTrackFirst[firstString[i]]; !isOk {
			listTrackFirst[firstString[i]] = 0
		}
		listTrackFirst[firstString[i]] += 1
	}
	for i := 0; i < len(secondString); i++ {
		if _, isOk := listTrackSecond[secondString[i]]; !isOk {
			listTrackSecond[secondString[i]] = 0
		}
		listTrackSecond[secondString[i]] += 1
	}

	for indexFirst, valFirst := range listTrackFirst {
		if val, isOk := listTrackSecond[indexFirst]; isOk {
			if valFirst <= val {
				totalEqual += valFirst
			} else {
				totalEqual += val
			}
		}
	}

	return len(firstString) - totalEqual
}
