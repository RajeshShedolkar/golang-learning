package main

import (
	"fmt"
)

func MaxSumOfSubArr(arr []int, size int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Empty Arr")
	}

	currSum := 0
	for _, v := range arr[:size] {
		currSum += v
	}
	k := size
	maxSum := currSum
	for k < len(arr) {
		currSum -= arr[k-size]
		currSum += arr[k]
		if maxSum < currSum {
			maxSum = currSum
		}
		k += 1
	}
	return maxSum, nil
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	size := 3
	fmt.Println(MaxSumOfSubArr(arr, size))

}
