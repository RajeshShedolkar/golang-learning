// The panic function to cause a run time error. 
// We can handle a run-time panic with the built-in recover function. 
// recover stops the panic and returns the value that was passed to the call 
// to panic.

// A panic generally indicates a programmer error 
// (for example attempting to access an index of an array 
// 	that's out of bounds, forgetting to initialize a map, etc.)


package main
import "fmt"

func main(){
	// panic("PANIC!!!")
	// str := recover()
	// fmt.Println("In Main function", str)
	defer func (){
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC!!!")
}