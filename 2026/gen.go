package main

import "fmt"

func Sum(arr []int) int {
	s := 0
	for i, v := range arr {
		fmt.Println("loop number -->", i)
		s += v
	}
	fmt.Println(s)
	return s
}

func MaxEle(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func ReverseArray(arr *[]int) {
	start, end := 0, len(*arr)-1
	mid := len(*arr) / 2
	fmt.Println(start, end, mid)
	for i := 0; i < mid; i++ {
		(*arr)[start], (*arr)[end] = swap((*arr)[start], (*arr)[end])
		start += 1
		end -= 1
	}
	fmt.Println(*arr)
}

func swap(a int, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func ReverseArray2(arr []int) {
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start += 1
		end -= 1
	}
}

func CheckArraySorted(arr []int)bool{
	first := 0
	for first < len(arr)-1{
		if arr[first] > arr[first+1]{
			return false
		}
		first+=1
	}
	return true
}
