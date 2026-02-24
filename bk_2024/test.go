package main
import "fmt"

// this is a comment
func main(){
	// hello()
	// test()
	// cal()
	variables()
}

func hello(){
	fmt.Println("Hello World")
	fmt.Println("1+1:", 1+1)
	fmt.Println(`Hello World 
well that's not a pbm
2`)
}

func test(){
	fmt.Println("len:", len("Heloo Worlds"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}

func cal(){
	fmt.Println(32132*42452)
}

func variables(){
	var x string  = "Hello World"
	fmt.Println(x)
	x+=" test"
	fmt.Println(x)
	var y string
	y = "new world"
	fmt.Println(y)
	y = x
	fmt.Println(y)
	x = "new x"
	fmt.Println(x, y)
	z := "New Z World"
	fmt.Println(z, len(z))

	w := 5
	fmt.Println(w)

	
}