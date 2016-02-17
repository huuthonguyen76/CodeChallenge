// Challenge Url: https://www.hackerrank.com/challenges/the-love-letter-mystery

package main

import (
	"fmt"
	"math"
)

type sortRunes []rune

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
		fmt.Println(totalReduce(cases[i]))
	}
}

func round(num float64) int {
	if num > math.Floor(num) {
		return int(math.Floor(num)) + 1
	}
	return int(math.Floor(num))
}

func totalReduce(text string) int {
	var reduceTimes = 0
	for i := 0; i < round(float64(len(text))/2); i++ {
		if text[i] != text[len(text)-i-1] {
			reduceTimes += int(math.Abs(float64(int(text[i]) - int(text[len(text)-i-1]))))
		}
	}
	return reduceTimes
}
