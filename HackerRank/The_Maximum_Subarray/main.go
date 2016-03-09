package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totalCases int
	var currentQuantities int
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	totalCases = ReadInt(s)

	for i := 0; i < totalCases; i++ {
		currentQuantities = ReadInt(s)
		currentCase := ReadInts(s, currentQuantities)
		fmt.Println(maximum(currentCase))
	}
}

func ReadInt(s *bufio.Scanner) int {
	s.Split(bufio.ScanWords)
	s.Scan()
	result, _ := strconv.Atoi(s.Text())

	return result
}

func ReadInts(s *bufio.Scanner, total int) []int64 {
	s.Split(bufio.ScanWords)
	a := make([]int64, total)
	for i := 0; i < total; i++ {
		s.Scan()
		a[i], _ = strconv.ParseInt(s.Text(), 10, 64)
	}
	return a
}

func maximum(list []int64) (int64, int64) {
	var isNegative = true
	for i := 0; i < len(list); i++ {
		if list[i] >= 0 {
			isNegative = false
			break
		}
	}

	if isNegative == true {
		var currentNum int64 = list[0]
		for i := 0; i < len(list); i++ {
			if currentNum < list[i] {
				currentNum = list[i]
			}
		}
		return currentNum, currentNum
	}

	var totalNon int64 = 0
	var total int64 = 0
	var currentTotal int64 = 0
	var listLen = len(list)

	for i := 0; i < listLen; i++ {
		if list[i] > 0 {
			totalNon += list[i]
		}
		currentTotal += list[i]

		if currentTotal < 0 {
			currentTotal = 0
		}
		if total < currentTotal {
			total = currentTotal
		}
	}

	return total, totalNon
}
