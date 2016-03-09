package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var a int
	var b int
	var c int
	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		fmt.Println(closestNumber(power(a, b), c))
	}
}

func power(num int, times int) int {
	var result = 1
	for i := 0; i < times; i++ {
		result *= num
	}
	return result
}

func closestNumber(num int, mod int) int {
	var excess = num % mod
	if excess > mod/2 {
		return num + (mod - excess)
	} else if excess < mod/2 || excess == mod/2 {
		return num - excess
	}
	return 0
}
