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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverseArr(arr))
}
