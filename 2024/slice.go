package main

import (
	"fmt"
)

func main (){
	// var x []float64;
	y := make([]float64, 5);
	fmt.Println(y);
	x := make([]float64, 5, 10);
	
	x[4]=1
	fmt.Println(x)

	arr := [5]float64{1, 2, 3,4 , 5};
	z := arr[1:4];
	fmt.Println(z);
	silce1 := append(x, 89, 56);
	fmt.Println(silce1)

}