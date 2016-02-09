// Challenge Url: https://www.hackerrank.com/challenges/pangrams

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var listTrack = make(map[uint8]bool)

	bio := bufio.NewReader(os.Stdin)
	currentScanString, _, _ := bio.ReadLine()
	sample := strings.ToLower(string(currentScanString))

	for i := 0; i < len(sample); i++ {
		if _, isOk := listTrack[sample[i]]; !isOk && sample[i] != 32 {
			listTrack[sample[i]] = true
		}
	}

	if len(listTrack) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
