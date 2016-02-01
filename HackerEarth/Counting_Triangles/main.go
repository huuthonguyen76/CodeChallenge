// Challenge Url: https://www.hackerearth.com/problem/algorithm/counting-triangles/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var cases = make(map[string]int)
	var totalCases int
	var results = 0

	bio := bufio.NewReader(os.Stdin)
	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		currentString, _, _ := bio.ReadLine()
		currentParts := strings.Split(string(currentString), " ")
		tempParts := make([]int, 0)

		// Convert string to increase number sequence
		for index := 0; index < len(currentParts); index++ {
			a, _ := strconv.Atoi(currentParts[index])
			tempParts = append(tempParts, a)
		}
		sort.Ints(tempParts)
		concatedString := convertNumsToString(tempParts)

		// Add the times string appear
		if _, ok := cases[concatedString]; ok {
			cases[concatedString] += 1
		} else {
			cases[concatedString] = 1
		}
	}

	// Count the result
	for _, element := range cases {
		if element == 1 {
			results += 1
		}
	}

	fmt.Print(results)
}

func convertNumsToString(list []int) string {
	var stringConverted string
	for index := 0; index < len(list); index++ {
		stringConverted += strconv.Itoa(list[index])
	}
	return stringConverted
}
