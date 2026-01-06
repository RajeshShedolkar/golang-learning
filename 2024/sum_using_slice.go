package main

import "fmt"

// By using ... before the type name of the last parameter
// you can indicate that it takes zero or more of those parameters.
// In this case we take zero or more ints.

func main() {
	a := []int{1, 2, 3, 44, 5, 6}
	sum(a[:3]...)
	fmt.Println(half(2))
	fmt.Println("Max Value:", max(a...))

	oddGen := makeOddGenerator()

	fmt.Println("oddGen:", oddGen())
	fmt.Println("oddGen:", oddGen())
	fmt.Println("oddGen:", oddGen())

	fmt.Println("Fib of 4:", fib(4))
}

func sum(args ...int) {
	s := 0
	for _, value := range args {
		s += value
	}
	fmt.Println(s)
}

func half(a int) bool {
	if (a/2)%2 == 0 {
		return true
	} else {
		return false
	}
}

func max(args ...int) int {
	max_value := args[0]
	for _, value := range args {
		if value > max_value {
			max_value = value
		}
	}
	return max_value
}

func makeOddGenerator() func() int {
	i := 1
	return func() int {
		ret := i
		i += 2
		return ret
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
