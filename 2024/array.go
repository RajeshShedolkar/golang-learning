package main
import "fmt"

func main(){
	var a [5]float32
	var x string
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("a", a, x+" what")
	var total float32 = 0
	for i:=0; i< len(a);i++{
		fmt.Println(a[i])
		total+=a[i]
	}
	fmt.Println(total/float32(len(a)))

	array()
}


func array(){
	var A [5]float64 
	A = [5]float64 {1, 2, 3, 4, 5}
	fmt.Println(A)
	for _, value := range A{
		fmt.Println(value)
	}
	B:=[3]float64{
		1,
		2,
		3,
	}
	fmt.Println(B)
}