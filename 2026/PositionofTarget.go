package main

import "fmt"

func positionOfTarget(arr []int, target int) []int {
	ind, err := getTarget(arr, target)
	start, end := ind, ind
	if err != nil {
		return []int{-1, -1}
	}

	for start-1 >= 0 {
		if arr[start-1] == arr[ind] {
			start -= 1
		} else {
			break
		}
	}

	for end+1 < len(arr) {
		if arr[end+1] == arr[ind] {
			end += 1
		} else {
			break
		}
	}

	return []int{start, end}
}

func positionOfTarget2(arr []int, target int) []int {
	ind, err := getTarget(arr, target)
	if err != nil {
		return []int{-1, -1}
	}

	min_ind := getMinInd(arr, ind)
	max_ind := getMaxInd(arr, ind)

	return []int{min_ind, max_ind}

}

func getMinInd(arr []int, ind int) int {

	return ind
}

func getMaxInd(arr []int, ind int) int {

	return ind
}

func getTarget(arr []int, target int) (int, error) {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid, nil
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1, fmt.Errorf("No Target found")

}

func main() {
	arr := []int{1, 2, 2, 2, 3, 4, 5}
	target := 2
	fmt.Println("Input arr: ", arr, "Target: ", target)
	fmt.Println(positionOfTarget(arr, target))

}
