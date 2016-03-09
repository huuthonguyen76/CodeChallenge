package main

import (
	"fmt"
	"math"
)

func main() {
	var totalCases int
	fmt.Scan(&totalCases)

	var currentStop1,
		currentStop2 int

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentStop1)
		fmt.Scan(&currentStop2)
		fmt.Println(countSquare(currentStop1, currentStop2))
	}
}

func countSquare(stop1 int, stop2 int) int {
	var start int
	var end int
	for i := stop1 - 1; i >= 1; i-- {
		if isRound(math.Sqrt(float64(i))) {
			start = int(math.Sqrt(float64(i)))
			break
		}
	}
	for i := stop2; i >= 1; i-- {
		if isRound(math.Sqrt(float64(i))) {
			end = int(math.Sqrt(float64(i)))
			break
		}
	}
	return end - start
}

func isRound(num float64) bool {
	if math.Floor(num) < num {
		return false
	}
	return true
}
