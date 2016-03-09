package main

import (
	"bufio"
	"fmt"
	_ "os"
	"strconv"
)

func main() {
	// var totalCases int
	// var currentQuantities int
	// var mod int64
	// fmt.Scan(&totalCases)

	// s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	// for i := 0; i < totalCases; i++ {
	// 	fmt.Scan(&currentQuantities)
	// 	fmt.Scan(&mod)

	// 	tokens := ReadInts(s, currentQuantities)
	// 	fmt.Println(maximiseSum(tokens, mod))
	// }
	var abc = []int64{1, 2, 3, 6, 9}
	fmt.Print(binarySearch(abc, 7))
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
