// Challenge Url: https://www.hackerrank.com/challenges/connected-cell-in-a-grid

package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

var currentRegion = 2
var n int
var m int

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	var matrix [][]int
	var currentInt int
	var result = 0

	for i := 0; i < n; i++ {
		currentRow := make([]int, 0)
		for j := 0; j < m; j++ {
			fmt.Scan(&currentInt)
			currentRow = append(currentRow, currentInt)
		}
		matrix = append(matrix, currentRow)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				currentResult := trackRegion(matrix, point{i, j})
				if result < currentResult {
					result = currentResult
				}
			}
		}
	}

	fmt.Println(result)
}

func trackRegion(matrix [][]int, coordinate point) int {
	var stack []point
	stack = append(stack, coordinate)
	result := 0
	for {
		if len(stack) == 0 {
			break
		}
		result++
		currentPoint := stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		matrix[currentPoint.x][currentPoint.y] = currentRegion

		for i := currentPoint.x - 1; i < currentPoint.x+2; i++ {
			for j := currentPoint.y - 1; j < currentPoint.y+2; j++ {
				if i < 0 || i > n-1 || j < 0 || j > m-1 {
					continue
				}
				if matrix[i][j] == 1 {
					matrix[i][j] = currentRegion
					stack = append(stack, point{i, j})
				}
			}
		}
	}
	return result
}
