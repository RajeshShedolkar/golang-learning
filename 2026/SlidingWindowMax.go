package main

import "fmt"

func pushDecresing(q []int, val int) []int {
	for len(q) > 0 && q[len(q)-1] < val {
		q = q[:len(q)-1]
	}
	q = append(q, val)
	return q
}

func pushDecresingIndex(arr []int, q []int, index int) []int {
	for len(q) > 0 && arr[q[len(q)-1]] < arr[index] {
		q = q[:len(q)-1]
	}
	q = append(q, index)
	return q
}

func maxSlidingWindow(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	result := make([]int, 0)
	que := make([]int, 0, len(arr))
	que1 := make([]int, 0, len(arr))

	s := 0
	for i, v := range arr {
		que1 = pushDecresing(que1, v)
		que = pushDecresingIndex(arr, que, i)
		fmt.Println("Iteration: ", i, "Queue: ", que, que1)
		if i-s+1 == k {
			result = append(result, arr[que[0]])
			s+=1
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}
