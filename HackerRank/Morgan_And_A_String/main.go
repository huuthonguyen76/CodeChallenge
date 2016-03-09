// Challenge Url: https://www.hackerrank.com/challenges/morgan-and-a-string

package main

import (
	"fmt"
)

func main() {
	var totalCases int
	var currentString1 string
	var currentString2 string
	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentString1)
		fmt.Scan(&currentString2)

		fmt.Println(out(currentString1, currentString2))
	}
}

func getRunes(text string) []byte {
	var result = make([]byte, 0)
	for i := 0; i < len(text); i++ {
		result = append(result, text[i])
	}
	return result
}

func out(text1, text2 string) string {
	rune1 := getRunes(text1)
	rune2 := getRunes(text2)
	result := make([]byte, 0)
	var step1 = 0
	var step2 = 0

	var len1 = len(rune1)
	var len2 = len(rune2)

	for {
		if step1 == len1 || step2 == len2 {
			break
		} else {
			if rune1[step1] <= rune2[step2] {
				result = append(result, rune1[step1])
				step1++
			} else if rune1[step1] > rune2[step2] {
				result = append(result, rune2[step2])
				step2++
			}
		}
	}
	for i := step1; i < len1; i++ {
		result = append(result, rune1[i])
	}
	for i := step2; i < len2; i++ {
		result = append(result, rune2[i])
	}
	return string(result)
}
