package main

import "fmt"

func BinarySearch(arr []int, target int) (int, error) {

	start, end := 0, len(arr)-1
// 1 2 3 4
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid, nil
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1, fmt.Errorf("No Element found")
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(BinarySearch(arr, 1))
}
