// Challenge Url: https://www.hackerearth.com/problem/algorithm/karan-and-strings/
package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var currentString string
	fmt.Scan(&totalCases)

	var cases = make([]string, 0)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentString)
		cases = append(cases, currentString)
	}

	var results = removeDuplicate(cases)
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}

}

func removeDuplicate(cases []string) []string {
	var processedCases = make([]string, 0)
	var currentChr uint8
	var currentCase string
	for i := 0; i < len(cases); i++ {
		currentChr = cases[i][0]
		currentCase = string(currentChr)
		for j := 0; j < len(cases[i]); j++ {
			if currentChr != cases[i][j] {
				currentCase += string(cases[i][j])
				currentChr = cases[i][j]
			}
		}
		processedCases = append(processedCases, currentCase)
	}
	return processedCases
}
