package main

import "fmt"

func main(){
	x := []float64{1, 2, 3, 4, 5}
	var y [3]float64;
	y = [3]float64{1,2,3};
	for i, value := range x{
		fmt.Println(i, value)
	}
	n:=2
	z := make([]float64, n)
	z = []float64{8, 9}
	fmt.Println(x, y, z)
	test(4)
	test_map()
	arr := []int{2, 4, 5, 2, 7}
	// fmt.Println(arr)
	find_duplicates(arr)
	smallest_num()
	var w int = 2;
	fmt.Println(w)

}

func test(n int){
	z := make([]float64, n);
	for i:=0; i<n; i++{
		z[i]=float64(i*2)
	}
	fmt.Println("From test function with size of array and value", n, z)
}

func test_map(){
	// var x map[string]string
	x := make(map[string]string)
	x["one"] = "1"
	x["two"] = "02"
	fmt.Println(x)
	y := x["two"]
	fmt.Println(y)
}

func find_duplicates(arr []int){
	d := make(map[int][]int)
	fmt.Println(arr)
	for _, value := range arr{
		fmt.Println(value, d[value])
		name, ok := d[value]
		// fmt.Println("This is extraction", name, ok)
		if !ok{
			d[value] = append(d[value], value)
		}else{
			fmt.Println("This is the duplicate", name,value)
		}
		// fmt.Println(ok)
		
	} 

// 	d1 := make(map[string]map[string]string)
// 	d1 = {"v": {
// 		"inner_test": "1",
// 		"inner_test2": "2"
// 	}
// }
// 	fmt.Println(d1)
}


func smallest_num(){
	arr := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17,}
	fmt.Println(arr)
	min_val:=arr[0]
	for _, value := range arr{
		if value < min_val{
			min_val=value
		}
	}
	fmt.Println("The min value:", min_val)
}