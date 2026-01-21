package main

func TwoSumPbm(arr []int, target int) []int {

	hash := make(map[int]int)

	for index, curr_ele := range arr {
		diff := target - curr_ele
		if i, ok := hash[diff]; ok {
			return []int{i, index}
		}
		hash[curr_ele] = index
	}
	return nil

}
