// Challenge Url: https://www.hackerrank.com/challenges/sherlock-and-array

package main

import (
	"fmt"
)

func main() {
	var totalCase int
	var currentTotalFrequency int
	var currentNum int
	var cases [][]int

	fmt.Scan(&totalCase)

	for i := 0; i < totalCase; i++ {
		fmt.Scan(&currentTotalFrequency)
		currentNums := make([]int, 0)
		for j := 0; j < currentTotalFrequency; j++ {
			fmt.Scan(&currentNum)
			currentNums = append(currentNums, currentNum)
		}
		cases = append(cases, currentNums)
	}

	for i := 0; i < len(cases); i++ {
		if len(cases[i]) == 1 || isRight(cases[i]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isRight(nums []int) bool {
	currentLeft := nums[0]
	currentRight := countTotal(nums[2:])

	if currentLeft == currentRight {
		return true
	}

	for j := 1; j+1 < len(nums); j++ {
		currentLeft += nums[j]
		currentRight -= nums[j+1]

		if currentLeft == currentRight {
			return true
		}

	}
	return false
}

func countTotal(nums []int) int {
	var total = 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// 1 2 3 3 1 5
