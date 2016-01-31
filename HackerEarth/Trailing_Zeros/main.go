// Challenge Url: https://www.hackerearth.com/problem/algorithm/trailing-zeros/
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Print(trailingZero(number))
}

func trailingZero(number int) int {
	var totalFive = 0
	for i := 1; i <= number; i++ {
		if i%5 == 0 {
			totalFive += fiveFactor(i)
		}
	}
	return totalFive
}

func fiveFactor(number int) int {
	var total = 0
	for number%5 == 0 {
		total += 1
		number = number / 5
	}
	return total
}
