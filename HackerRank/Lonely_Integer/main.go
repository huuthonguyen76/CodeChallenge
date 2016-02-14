// Challenge Url: https://www.hackerrank.com/challenges/lonely-integer

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var currentNum int
	fmt.Scan(&totalCases)

	var listTracker = make(map[int]bool)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentNum)
		if _, isOk := listTracker[currentNum]; isOk {
			listTracker[currentNum] = true
		} else {
			listTracker[currentNum] = false
		}
	}

	for index, val := range listTracker {
		if val == false {
			fmt.Print(index, " ")
		}
	}
}
