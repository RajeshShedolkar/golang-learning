package main

import "fmt"

func getMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxVal := arr[0]
	for _, v := range arr {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func getSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func isFeasible(boards []int, k int, T int) bool {
	// write code here
	if len(boards) == 0 {
		return true
	}
	if k <= 0 {
		return false
	}
	painters := 1
	curr := 0

	for _, b := range boards {
		if b > T {
			return false
		}
		if curr+b <= T {
			curr += b
		} else {
			painters++
			curr = b

			if painters > k {
				return false
			}
		}
	}
	return true
}

func MinTime(boards []int, k int) int {
	minTime := getMax(boards)
	maxTime := getSum(boards)
	T := 0
	for minTime <= maxTime {
		T = (minTime + maxTime) / 2
		if isFeasible(boards, k, T) {
			maxTime = T - 1
		} else {
			minTime = T + 1
		}
	}
	return minTime
}

func main() {
	tests := []struct {
		boards []int
		k      int
		want   int
	}{
		{[]int{10, 20, 30, 40}, 2, 60}, // split: [10,20,30] | [40]
		{[]int{10, 20, 30, 40}, 3, 40}, // [10,20] | [30] | [40]
		{[]int{5, 5, 5, 5}, 2, 10},     // [5,5] | [5,5]
		{[]int{7, 2, 5, 10, 8}, 2, 18}, // [7,2,5] | [10,8]
		{[]int{7, 2, 5, 10, 8}, 3, 14}, // [7,2,5] | [10] | [8]
		{[]int{1}, 3, 1},               // single board
	}

	for i, tc := range tests {
		got := MinTime(tc.boards, tc.k)
		fmt.Printf("Test #%d: got=%d want=%d -> %v\n", i+1, got, tc.want, got == tc.want)
	}
}
