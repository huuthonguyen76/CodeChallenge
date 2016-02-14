// Challenge Url: https://www.hackerrank.com/challenges/make-it-anagram

package main

import (
	"fmt"
)

func main() {
	var listTrackFirst = make(map[uint8]int)
	var listTrackSecond = make(map[uint8]int)

	var firstString string
	var secondString string

	fmt.Scan(&firstString)
	fmt.Scan(&secondString)

	var totalDeletion = 0

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
		if val, isOk := listTrackSecond[indexFirst]; !isOk {
			totalDeletion += valFirst
		} else {
			odd := valFirst - val
			if odd < 0 {
				odd = odd * -1
			}
			totalDeletion += odd
			listTrackSecond[indexFirst] = -1
		}
	}

	for _, val := range listTrackSecond {
		if val != -1 {
			totalDeletion += val
		}
	}

	fmt.Print(totalDeletion)

}
