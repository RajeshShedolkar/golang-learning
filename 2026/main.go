package main

import (
	"fmt"
	"reflect"
)

//import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s = append(s, 5)
	// fmt.Println(Sum(s))
	// fmt.Println(TwoSumPbm(s, 6))
	// fmt.Println(MaxEle(s))

	//fmt.Println()
	// ReverseArray(&s)
	// fmt.Println(s)
	//ReverseArray2(s)
	fmt.Println(CheckArraySorted([]int{1, 2, 3, 4}))
	fmt.Println(CheckArraySorted([]int{3, 1, 2, 3}))
	fmt.Println(s)
	s1 := [][]int{
		{1, 2, 3, 4, 5},   // sorted
		{5, 4, 3, 2, 1},   // reverse sorted
		{2, 2, 2, 2},      // all equal
		{1},               // single element
		{},                // empty array
		{-5, -2, -10, -1}, // negative numbers
		{1, 3, 2, 4, 5},   // not sorted
		{5, 3, 4},
	}
	for _, curr_arr := range s1 {
		fmt.Println(SecondLargestElement2(curr_arr))
	}
	TestCaseKSum()
}

func TestCaseKSum() {
	tests := []struct {
		name     string
		arr      []int
		k        int
		expected []int
	}{
		// 1. Basic cases
		{"basic_1", []int{1, 2, 3, 4, 5}, 3, []int{6, 9, 12}},
		{"k_equals_1", []int{10, 20, 30}, 1, []int{10, 20, 30}},
		{"all_same", []int{5, 5, 5, 5}, 2, []int{10, 10, 10}},

		// 2. Edge cases
		{"empty_array", []int{}, 3, []int{}},
		{"k_greater_than_len", []int{1, 2}, 3, []int{}},
		{"k_zero", []int{1, 2, 3}, 0, []int{}},
		{"single_element", []int{7}, 1, []int{7}},

		// 3. Negative numbers
		{"all_negative", []int{-1, -2, -3, -4}, 2, []int{-3, -5, -7}},
		{"mixed_signs", []int{1, -1, 2, -2, 3}, 2, []int{0, 1, 0, 1}},

		// 4. Mixed values
		{"mixed_values", []int{100, -50, 200, -100, 50}, 3, []int{250, 50, 150}},

		// 5. All zeros
		{"all_zeros", []int{0, 0, 0, 0}, 2, []int{0, 0, 0}},
	}

	for _, tc := range tests {
		// 👉 Replace `SumOfSubarraysOfSizeK` with your function name
		result := GetkSizeSum2(tc.arr, tc.k)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ %s passed\n", tc.name)
		} else {
			fmt.Printf("❌ %s failed | expected=%v got=%v\n",
				tc.name, tc.expected, result)
		}
	}
}
