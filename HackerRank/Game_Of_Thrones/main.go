// Challenge Url: https://www.hackerrank.com/challenges/game-of-thrones

package main

import (
	"fmt"
)

func main() {
	var listTrack = make(map[uint8]int)
	var oddCount int = 0
	var sample string
	fmt.Scan(&sample)

	for i := 0; i < len(sample); i++ {
		if _, isOk := listTrack[sample[i]]; !isOk {
			listTrack[sample[i]] = 0
		}
		listTrack[sample[i]] += 1
	}

	for _, val := range listTrack {
		if val%2 != 0 {
			oddCount += 1
		}
		if oddCount > 1 {
			fmt.Print("NO")
			break
		}
	}
	if oddCount <= 1 {
		fmt.Print("YES")
	}
}
