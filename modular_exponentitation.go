package main

import (
	"fmt"
)

func main() {
	var a = 2
	var b = 4
	var c = 6
	fmt.Println(modular(a, b, c))
}

func modular(a int, b int, c int) int {
	var ans = 1
	for b != 0 {
		if b%2 == 1 {
			ans = (ans * a) % c
		}
		a = (a * a) % c
		b /= 2
	}
	return ans
}
