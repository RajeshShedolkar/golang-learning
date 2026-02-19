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

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	maxLen := 0

	for _, num := range nums {

		if !set[num-1] {

			current := num
			currLen := 1

			for set[current+1] {
				current++
				currLen++
			}

			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}

	return maxLen
}

func main() {
	arr := []int{7, 3, 4, 5, 1, 2, 0}
	fmt.Println(longestConsecutive(arr))
}
