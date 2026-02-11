package main

import (
	"errors"
	"fmt"
)

func MaxSumOfSubArr(arr []int, size int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}

	if size <= 0 || size > len(arr) {
		return 0, errors.New("invalid window size")
	}

	currSum := 0
	for _, v := range arr[:size] {
		currSum += v
	}

	maxSum := currSum

	for k := size; k < len(arr); k++ {
		currSum -= arr[k-size]
		currSum += arr[k]

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum, nil
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	size := 3

	ans, err := MaxSumOfSubArr(arr, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Max Sum:", ans)
}
