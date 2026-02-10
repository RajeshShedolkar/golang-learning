package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	deque := make([]int, 0) // stores indices
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		// 1. Remove indices out of window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 2. Remove smaller elements from back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 3. Add current index
		deque = append(deque, i)

		// 4. Add max to result once window is formed
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}

