package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findInsert(list []int64, value int64) int {
	var start = 0
	var end = len(list)
	var midpoint = 0

	for start <= end {
		midpoint = (start + end) / 2
		if list[midpoint] == value {
			return midpoint
		} else if list[midpoint] < value {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}
	}

	return midpoint
}

func findLower(list []int64, value int64) int {
	var start = 0
	var end = len(list)
	var midpoint = 0
	var result = -1

	if value > list[end-1] {
		return end - 1
	}
	if value < list[0] {
		return -1
	}

	for start <= end {
		midpoint = (start + end) / 2
		if list[midpoint] == value {

			for i := midpoint; i >= 0; i-- {
				if list[i] < list[midpoint] {
					result = i
					break
				}
			}
			return result
		} else if list[midpoint] < value {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}
	}

	for i := midpoint; i >= 0; i-- {
		if list[i] < value {
			result = i
			break
		}
	}

	return result
}

func insert(list *[]int64, value int64) {
	var _list = *list

	if len(_list) == 0 {
		_list = append(_list, value)
	} else if value <= _list[0] {
		_list = append([]int64{value}, _list...)
	} else if value >= _list[len(_list)-1] {
		_list = append(_list, value)
	} else {
		var position = findInsert(_list, value)
		if _list[position] == value {
			_list = append(_list[:position+1], _list[position:]...)
		} else if _list[position] > value {
			_list = append(_list[:position], append([]int64{value}, _list[position:]...)...)
		} else {
			_list = append(_list[:position+1], append([]int64{value}, _list[position+1:]...)...)
		}
	}
	*list = _list
}

func ReadInts(s *bufio.Scanner, total int) []int {
	s.Split(bufio.ScanWords)
	a := make([]int, total)
	for i := 0; i < total; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	return a
}

func ReadInt(s *bufio.Scanner) int {
	s.Split(bufio.ScanWords)
	s.Scan()

	result, _ := strconv.Atoi(s.Text())
	return result
}

func main() {
	var total int64 = 0
	var currentQueries int64 = 0
	var trackList = make(map[int64]int64)
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	totalNums := ReadInt(s)
	nums := ReadInts(s, totalNums)

	totalQueries := ReadInt(s)
	queries := ReadInts(s, totalQueries)

	sort.Ints(nums)

	var list = make([]int64, totalNums)
	for i := 0; i < totalNums; i++ {
		list[i] = int64(nums[i])
	}

	for i := 0; i < totalNums; i++ {
		total += list[i]
		trackList[list[i]] = total
	}

	// -3 -1 2
	for i := 0; i < totalQueries; i++ {
		currentQueries += int64(queries[i])
		lower := int64(findLower(list, currentQueries*-1))
		if lower == -1 {
			fmt.Println(trackList[list[totalNums-1]] + currentQueries*(int64(totalNums)))
		} else {
			var a1 = ((lower+1)*currentQueries + trackList[list[lower]]) * -1
			var a2 = trackList[list[totalNums-1]] - trackList[list[lower]] + (int64(totalNums)-1-lower)*currentQueries
			fmt.Println(a1 + a2)
		}
	}
}
