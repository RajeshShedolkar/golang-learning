package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s = append(s, 5)
	fmt.Println(Sum(s))
	fmt.Println(TwoSumPbm(s, 6))
	fmt.Println(MaxEle(s))
}
