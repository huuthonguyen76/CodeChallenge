// Challenge Url: https://www.hackerrank.com/challenges/staircase

package main

import (
	"fmt"
)

func main() {
	var height int
	fmt.Scan(&height)
	height += 1

	for i := 1; i < height; i++ {
		for j := 1; j < height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
