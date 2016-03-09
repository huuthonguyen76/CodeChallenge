// Challenge Url: https://www.hackerrank.com/challenges/utopian-tree

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	fmt.Scan(&totalCases)

	var currentCycle int
	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentCycle)
		fmt.Println(heightTree(currentCycle))
	}
}

func heightTree(cycle int) int {
	var height = 1
	for i := 1; i <= cycle; i++ {
		if i%2 == 1 {
			height *= 2
		} else {
			height += 1
		}
	}
	return height
}
