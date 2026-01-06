package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	fmt.Println(time.TimeOnly)
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250)) 
		time.Sleep(time.Millisecond * amt) 
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(0)
	}
	var input string = "rajesh"
	fmt.Scanln(&input)
	time.Sleep(0)
	fmt.Println(input)
}
