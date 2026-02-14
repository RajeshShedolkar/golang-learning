package main

import "fmt"

func pushDecresing(q []int, val int) []int {
	for len(q) > 0 && q[len(q)-1] < val {
		q = q[:len(q)-1]
	}
	q = append(q, val)
	return q
}

func maxSlidingWindow(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	result := make([]int, 0)
	que := make([]int, len(arr), len(arr))

	for i, v := range arr {
		que = pushDecresing(que, v)
		fmt.Println("Iteration: ",i, "Queue: ", que)
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}
