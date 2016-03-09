package main

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

func main() {
	var matrix = make([][]int, 0)
	matrix = append(matrix, []int{1, 1, 1, 1, 1})
	matrix = append(matrix, []int{1, 1, 1, 1, 1})
	matrix = append(matrix, []int{1, 1, 3, 1, 1})
	matrix = append(matrix, []int{1, 1, 1, 1, 1})
	matrix = append(matrix, []int{1, 1, 1, 1, 1})
	matrix = append(matrix, []int{1, 2, 1, 1, 1})

}

func addBoundary(matrix *[][]int) {
	var newMatrix = make([][]int, 0)
	var length = len((*matrix)[0])
	var height = len
	var newMatrix = append()
}

func dfs(matrix [][]int, src vertex, dest vertex) {
	var stack = make([]vertex, 0)
	var currentVertex vertex
	stack = append(src)
	for {
		currentVertex = popStack(&stack)

	}
}

func popStack(stack *[]vertex) vertex {
	var stackLen = len(*stack)
	var result = (*stack)[stackLen-1]
	*stack = (*stack)[:stackLen-1]
	return result
}
