package main

func isSumExist(arr []int, target int) bool {

	start, end := 0, len(arr)-1

	for start < end {
		sum := arr[start] + arr[end]
		if sum == target {
			return true
		}
		if sum > target {
			end -= 1
		} else {
			start += 1
		}

	}
	return false
}
