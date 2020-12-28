package main

import (
	"fmt"
)

func main() {
	numMap := make(map[int]int)
	numMap[1] = 1
	numMap[12] = 2
	numMap[0] = 3
	numMap[20] = 4
	numMap[8] = 5
	numMap[16] = 6
	nextNum := 0
	for turnNum := 7; turnNum < 30000000; turnNum++ {
		if val, ok := numMap[nextNum]; ok {
			numMap[nextNum] = turnNum
			nextNum = turnNum - val
		} else {
			numMap[nextNum] = turnNum
			nextNum = 0
		}
	}
	fmt.Println(nextNum)
}
