package main

import "fmt"

func findMaxLenSubArray(arr []int, k int) (int, int) {
	if len(arr) <= 0 {
		return -1, -1
	}
	start, end := 0, 0
	size := 1
	for size <= len(arr) {
		currSizeSum := sumArrSize(arr, size)
		currStart, currEnd := 0, size
		for currEnd < len(arr) {
			if currSizeSum == k {
				if (currEnd - currStart) > (end - start) {
					start = currStart
					end = currEnd
				}
			}

			currStart += 1
			currEnd += 1
		}
		size += 1

	}
	return start, end
}

func sumArrSize(arr []int, size int) int {
	sumArr := 0
	for i := 0; i < size; i++ {
		sumArr += arr[i]
	}
	return sumArr
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	k := 3
	fmt.Println(findMaxLenSubArray(arr, k))
}
