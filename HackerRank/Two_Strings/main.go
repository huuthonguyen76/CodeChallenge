// Challenge Url: https://www.hackerrank.com/challenges/two-strings

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var firstString string
	var secondString string
	var results = make([]bool, 0)

	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&firstString)
		fmt.Scan(&secondString)
		results = append(results, haveSubString(firstString, secondString))
	}

	for i := 0; i < len(results); i++ {
		if results[i] == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func haveSubString(firstString string, secondString string) bool {
	var listTrackFirst = make(map[uint8]int)
	var listTrackSecond = make(map[uint8]int)

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

	for indexFirst, _ := range listTrackFirst {
		if _, isOk := listTrackSecond[indexFirst]; isOk {
			return true
		}
	}

	return false
}
