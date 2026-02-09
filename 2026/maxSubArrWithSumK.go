package main

func maxLenSubArr(arr []int, target int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	hash := make(map[int]int)
	currSum := 0
	maxLen := 0
	start, end := -1, -1

	for i, v := range arr {
		currSum += v

		if currSum == target {
			if i+1 > maxLen {
				maxLen = i + 1
				start = 0
				end = i
			}
		}

		if idx, ok := hash[currSum-target]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
				start = idx + 1
				end = i
			}
		}

		if _, exists := hash[currSum]; !exists {
			hash[currSum] = i
		}
	}

	return start, end
}
