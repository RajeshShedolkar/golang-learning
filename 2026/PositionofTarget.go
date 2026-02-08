package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{
		findFirst(nums, target),
		findLast(nums, target),
	}
}

func findFirst(nums []int, target int) int {
	start, end := 0, len(nums)-1
	res := -1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			res = mid
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}

func findLast(nums []int, target int) int {
	start, end := 0, len(nums)-1
	res := -1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			res = mid
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}


func main() {
	arr := []int{1, 2, 2, 2, 2, 2, 3, 4, 5}
	target := 2
	fmt.Println("Input arr: ", arr, "Target: ", target)
	fmt.Println(searchRange(arr, target))

}
