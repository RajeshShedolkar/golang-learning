package main
import "fmt"

type Circle1 struct{
	x float64
	y float64
	r float64
}
type Circle2 struct{x, y, r float64}

func main(){
	fmt.Println("Inside Main func")
	var c1 Circle1
	var c2 Circle2
	c := new(Circle1)
	c1.r, c2.r = 12.0, 123.0
	fmt.Println(c1, c2, *c)
}