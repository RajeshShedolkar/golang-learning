package main
import "fmt"

// clouser is the way to create function inside function


func main(){
	nextGen := evenGenrator()
	fmt.Println(nextGen())
	fmt.Println(nextGen())
	fmt.Println(nextGen())
	fmt.Println(nextGen())
	// add func
	addFunc()

	// nextEven := makeEvenGenerator() 
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())

	incrementFunc()

}

func evenGenrator() (func() uint){
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return ret
	}
}

func addFunc(){
	add := func(a , b int) int{
		return a+b
	}

	fmt.Println(add(1, 2))
}

func incrementFunc(){
	x := 0
	increment := func ()int{
		x++
		return x
	}
	fmt.Println("Increament:", increment())
	fmt.Println("Increament:", increment())
}