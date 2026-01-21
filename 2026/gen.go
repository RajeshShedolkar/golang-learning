package main

import (
	"fmt"
	"math"
)

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

func CheckArraySorted(arr []int) bool {
	first := 0
	for first < len(arr)-1 {
		if arr[first] > arr[first+1] {
			return false
		}
		first += 1
	}
	return true
}

func SecondLargestElement(arr []int) (int, bool) {
	fmt.Println("Input Array: ", arr)
	var maxEle, secondMaxEle int
	if len(arr) < 2 {
		return secondMaxEle, false
	}

	for _, curr_ele := range arr {
		if curr_ele > maxEle {
			maxEle = curr_ele
		}
	}

	for _, curr_ele := range arr {
		if curr_ele > secondMaxEle && curr_ele < maxEle {
			fmt.Println(maxEle, secondMaxEle)
			secondMaxEle = curr_ele
		}
	}

	if maxEle == secondMaxEle {
		return 0, false
	}
	return secondMaxEle, true
}

func SecondLargestElement2(arr []int) (int, bool) {
	fmt.Println("Input Array: ", arr)
	var maxEle, secondMaxEle int = math.MinInt, math.MinInt
	if len(arr) < 2 {
		return secondMaxEle, false
	}
	// maxEle, secondMaxEle
	for _, curr_ele := range arr {
		if curr_ele > maxEle {
			secondMaxEle = maxEle
			maxEle = curr_ele
		}
		if curr_ele > secondMaxEle && maxEle > curr_ele {
			secondMaxEle = curr_ele
		}

	}

	if maxEle == secondMaxEle {
		return 0, false
	}
	return secondMaxEle, true
}

// []int{1, 2, 3, 4, 5}
func GetkSizeSum(arr []int, k int) []int {
	if len(arr) == 0 || k > len(arr) || k == 0 {
		return []int{}
	}
	start, end := 0, len(arr)
	out := []int{0}
	for _, v := range arr[:k] {
		out[0] += v
	}

	for start+k < end {
		out = append(out, out[start]-arr[start]+arr[start+k])
		start += 1
	}
	return out
}

// []int{1, 2, 3, 4, 5}
func GetkSizeSum2(arr []int, k int) []int {
	n := len(arr)
	if len(arr) == 0 || k <= 0 || k > n {
		return []int{}
	}
	start, end := 0, len(arr)
	out := make([]int, 0, n-k)
	windowSum := 0
	for _, v := range arr[:k] {
		windowSum += v
	}
	out = append(out, windowSum)
	for start+k < end {
		currSum := out[start] - arr[start] + arr[start+k]
		out = append(out, currSum)
		start += 1
	}
	return out
}

// Arrays: Move all zeros to the end
// {1, 1, 0}
func MoveZerosToEnd(arr []int) []int{
	start, end := 0, len(arr)-1
	for start < end {
		if arr[start] != 0 {
			start += 1
			continue
		}
		if arr[end] == 0 {
			end -= 1
			continue
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start += 1
			end -= 1
		}
	}
	return arr
}

func MoveZerosToEndByPreservingNonZeroOrder(arr []int) []int{
	start, n := 0, len(arr)
	countNonZero := 0

	for i:=0;i< n; i++ {
		if arr[i] != 0{
			arr[start]=arr[i]
			countNonZero+=1
			start+=1
		}
	}
	for countNonZero < n{
		arr[countNonZero]=0
		countNonZero+=1
	}
	return arr
}
