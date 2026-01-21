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
