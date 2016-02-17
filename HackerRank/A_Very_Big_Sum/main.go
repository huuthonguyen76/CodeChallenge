// Challenge Url: https://www.hackerrank.com/challenges/a-very-big-sum

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var currentNum int64
	var total int64

	fmt.Scan(&totalCases)
	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentNum)
		total += currentNum
	}
	fmt.Println(total)
}
