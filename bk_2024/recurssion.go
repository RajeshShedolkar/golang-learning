package main
import "fmt"

func main(){
	fmt.Println(fact(5))
}

func fact(x uint) uint{
	if x==0{
		return 1
	}
	return fact(x-1)*x
}