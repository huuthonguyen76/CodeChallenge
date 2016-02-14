// Challenge Url: https://www.hackerrank.com/challenges/sherlock-and-anagrams

package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func main() {
	var totalCases int
	var currentString string
	var cases []string

	fmt.Scan(&totalCases)

	for i := 0; i < totalCases; i++ {
		fmt.Scan(&currentString)
		cases = append(cases, currentString)
	}

	for i := 0; i < len(cases); i++ {
		fmt.Println(findTotalAnagrams(cases[i]))
	}
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func findTotalAnagrams(text string) int {
	var totalAnagrams = 0
	var trackList = make(map[string]int)
	var currentString string

	for i := 0; i < len(text); i++ {
		for j := i + 1; j <= len(text); j++ {
			currentString = SortString(text[i:j])
			if _, isOk := trackList[currentString]; !isOk {
				trackList[currentString] = 0
			} else {
				trackList[currentString] += 1
				totalAnagrams += trackList[currentString]
			}
		}
	}
	return totalAnagrams
}
