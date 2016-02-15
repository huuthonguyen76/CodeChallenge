// Challenge Url: https://www.hackerrank.com/challenges/palindrome-index

package main

import (
	"fmt"
	"math"
)

func main() {
	var totalCases int
	var currentString string
	var cases []string

	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentString)
		cases = append(cases, currentString)
	}

	for i := 0; i < len(cases); i++ {
		fmt.Println(findPalndromeIndex(cases[i]))
	}

}

func findPalndromeIndex(text string) int {
	if isPalndrome(text) {
		return -1
	}

	for i := 0; i < len(text); i++ {
		if text[i] != text[len(text)-i-1] {
			if isPalndrome(text[0:i] + text[i+1:]) {
				return i
			}
		}
	}

	return 0
}

func isPalndrome(text string) bool {
	var loopTime = float64(len(text)) / 2
	var floorLoop = math.Floor(loopTime)
	var totalLoop int

	if floorLoop < loopTime {
		totalLoop = int(floorLoop + 1)
	} else {
		totalLoop = int(floorLoop)
	}

	for i := 0; i < totalLoop; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}
