package main

import (
	"fmt"
	"sort"
)

func longSeq(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	maxLen := 1
	currLen := 1

	for i := 1; i < len(arr); i++ {

		if arr[i] == arr[i-1] {
			continue
		}

		if arr[i]-arr[i-1] == 1 {
			currLen++
		} else {
			currLen = 1
		}

		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}


func main() {
	arr := []int{0, 1, 2,3, 2, 5}
	fmt.Println(longSeq(arr))
}
