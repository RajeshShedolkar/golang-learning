package main

import (
	"fmt"
	"math"
)

func MaxSumSubArr(arr []int) int {
	if len(arr) == 0 {
		return math.MinInt
	}

	currSum := 0
	maxSum := math.MinInt

	for _, currEle := range arr {
		currSum += currEle

		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSumSubArr(arr))
}
