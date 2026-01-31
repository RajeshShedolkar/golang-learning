package main

import "fmt"

func RemoveDuplicatesFromSortedArr(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	write := 0

	for read := 1; read < len(arr); read++ {
		if arr[read] != arr[write] {
			write++
			arr[write] = arr[read]
		}
	}
	return arr[:write+1]
}


func main() {
	arr := []int{1, 1, 2, 2, 3}
	// arr := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(RemoveDuplicatesFromSortedArr(arr))
}
