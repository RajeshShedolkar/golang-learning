package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxWater(arr []int) int {
	start, end := 0, len(arr)-1
	maxWater := 0
	for start <= end {
		currWater := min(arr[start], arr[end]) * (end - start)
		if maxWater < currWater {
			maxWater = currWater
		}
		if arr[start] < arr[end] {
			start += 1
		} else {
			end -= 1
		}

	}
	return maxWater
}

func main() {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{5, 5, 5, 5}, 15},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
		{[]int{1, 100, 1, 1, 100, 1}, 400},
	}

	for i, tc := range tests {
		result := MaxWater(tc.input)
		fmt.Printf("Test %d | Input: %v | Output: %d | Expected: %d\n",
			i+1, tc.input, result, tc.expect)
	}
}
