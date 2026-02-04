package main

import "fmt"

func findMaxLenSubArray(arr []int, k int) (int, int) {
	maxLen := 0
	startIdx, endIdx := -1, -1

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == k {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					startIdx = i
					endIdx = j
				}
			}
		}
	}
	return startIdx, endIdx
}


func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	k := 3
	fmt.Println(findMaxLenSubArray(arr, k))
}
