package main
import "fmt"

// defer is often used when resources need to be freed in some way. 
// For example when we open a file we need to make sure to close it later.

// This has 3 advantages: 
// (1) it keeps our Close call near our Open call so its easier to understand, 
// (2) if our function had multiple return statements (perhaps one in an if and one in an else) 
// 	Close will happen before both of them 
// (3) deferred functions are run even if a run-time panic occurs.


func main(){
	defer first()
	second()
	third()
	fourth()
}

func first(){
	fmt.Println("File close function")
}

func second(){
	fmt.Println("2nd close function")
}


func third(){
	fmt.Println("3rd close function")
}

func fourth(){
	fmt.Println("4th close function")
}