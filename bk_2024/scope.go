package main

import "fmt"

var x string = "Test"
var y uint8 = 10

func  main(){
	fmt.Println(x, y)
	f1()
	fmt.Println(x, y)
	forloop()
}

func f1(){
	x = "test"
	fmt.Println("from f1 function", x)
}

func forloop(){
	i:=1
	for i <= 10 {
		
		if i%2==0{
			fmt.Println(i)
		}else{

		}
		switch i {
		case 1: {
			a := i
			fmt.Println("one", a)
		}
		case 2: fmt.Println("two", i)
	default: fmt.Println("Rest")
		}

		i+=1
	}


	// for i:=1; i<10; i++{
	// 	fmt.Println(i)
	// }
}