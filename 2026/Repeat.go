package main

import "fmt"

func reverseArr(arr []int) []int {
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start += 1
		end -= 1
	}
	return arr
}

func isArrSorted(arr []int) bool {
	asc_flag := true
	dsc_flag := true
	for i := 1; i < len(arr); i++ {
		if asc_flag && arr[i] < arr[i-1] {
			asc_flag = false
		}
		if dsc_flag && arr[i] > arr[i-1] {
			dsc_flag = false
		}

		if !asc_flag && !dsc_flag {
			return false
		}
	}
	return asc_flag || dsc_flag
}

func moveAllZerosAtEnd(arr []int) []int {
	write := 0

	for read := 0; read < len(arr); read++ {
		if arr[read] != 0 {
			arr[write] = arr[read]
			write++
		}
	}

	for write < len(arr) {
		arr[write] = 0
		write++
	}

	return arr
}

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6}
	arr := []int{2, 1, 2, 3, 4}
	fmt.Println(reverseArr(arr))
	fmt.Println("isSorted: ", arr, isArrSorted(arr))
	a := []int{0, 1, 0, 3, 12}
	fmt.Println(moveAllZerosAtEnd(a))
}
