package main

import "fmt"

func pop(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	return arr[1:]
}
func popFromBack(arr []int)[]int{
	return arr[:len(arr)-1]
}

func maxSlidingWindow(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	result := make([]int, 0)
	// Write the Logic
	que := make([]int, len(arr), len(arr))
	for i, v := range arr {
		if len(que)==0{
			que = append(que, v)
		} 
		j := i
		for que[j-1]< v{
			
			que = popFromBack(que)
			que = append(que, v)
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}
