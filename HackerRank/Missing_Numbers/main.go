package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ReadInts(s *bufio.Scanner) []int {
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	return a
}

func main() {
	var result []int
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	tokens1 := ReadInts(s)
	tokens2 := ReadInts(s)

	var cases1 = make(map[int]int)
	var cases2 = make(map[int]int)

	var totalCases1 = len(tokens1)
	var totalCases2 = len(tokens2)

	for i := 0; i < totalCases1; i++ {
		cases1[tokens1[i]] += 1
	}

	for i := 0; i < totalCases2; i++ {
		cases2[tokens2[i]] += 1
	}

	for index1, val1 := range cases1 {
		if val2, isOk := cases2[index1]; !isOk {
			result = append(result, index1)
		} else {
			if val1 != val2 {
				result = append(result, index1)
			}
			delete(cases2, index1)
		}
	}

	for index, _ := range cases2 {
		result = append(result, index)
	}
	sort.Ints(result)
	lenResult := len(result)
	for i := 0; i < lenResult; i++ {
		fmt.Print(result[i], " ")
	}
}
