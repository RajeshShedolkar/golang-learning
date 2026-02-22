package main

func getMax(arr []int) int {
	if len(arr)==0{
		return 0
	}
	maxVal := arr[0]
	for _, v := range arr{
		if maxVal < v{
			maxVal = v
		}
	}
	return maxVal
}


func getSum(arr []int)int{
	sum := 0
	for _, v := range arr{
		sum += v
	}
	return sum
}


func isFeasible(boards []int, k int, T int)bool{

	return true
}

func MinTime(boards []int, k int)int{
	minTime := getMax(boards)
	maxTime := getSum(boards)
	T := 0
	for minTime < maxTime {
		T = (minTime + maxTime) /2 
		if isFeasible(boards, k, T){
			maxTime = T - 1
		}else{
			minTime = T + 1
		}
	}
	return T
}

func main(){

}