// Challenge Url: https://www.hackerrank.com/challenges/alternating-characters

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var cases = make([]string, 0)
	var currentScanString string
	var result []int

	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentScanString)
		cases = append(cases, currentScanString)
	}
	result = countDeletion(cases)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func countDeletion(cases []string) []int {
	var listDeletion []int
	var currentChr uint8
	var currentDeletion int

	for i := 0; i < len(cases); i++ {
		currentChr = cases[i][0]
		currentDeletion = 0

		for j := 1; j < len(cases[i]); j++ {
			if currentChr != cases[i][j] {
				currentChr = cases[i][j]
			} else {
				currentDeletion++
			}
		}
		listDeletion = append(listDeletion, currentDeletion)
	}
	return listDeletion
}
