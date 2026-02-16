package main
import (
	"fmt"
	"sort"
	"math"
)

func minDiff(arr []int, M int){
	sort.Ints(arr)
	start, end := 0, M-1
	minDiff := math.MaxInt

	for end < len(arr){
		start = end-M + 1
		if minDiff > arr[end]-arr[start]{
			minDiff = arr[end] - arr[start]
		}
		end += 1
	}

	fmt.Println("Distribution: ", minDiff)
}

func main(){
	arr := []int{3, 4, 1, 9, 56, 7, 9, 12}
	m := 5
	minDiff(arr, m)
}