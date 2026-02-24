package main

import "fmt"

func main() {
	x := 10
	fmt.Println("before", x)
	mod(&x)
	fmt.Println(x)
	fmt.Println("after", x)
	var y float64 = 10.0
	square(&y)
	fmt.Println(y)
	a := 1
	b := 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	z := new(int)
	*z = 33
	fmt.Println("z using new:", *z)
	change(z)
	fmt.Println("changed z:", *z)
}

func mod(x *int) {
	*x = 12
}

func square(x *float64) {
	*x = *x * *x
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func change(z *int) {
	*z = 21
}
