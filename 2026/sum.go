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
