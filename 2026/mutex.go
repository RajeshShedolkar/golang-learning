package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int

func addMoney(wg *sync.WaitGroup) {
	defer wg.Done()

	tmp := balance           // read
	time.Sleep(1 * time.Millisecond)
	tmp = tmp + 100          // modify
	time.Sleep(1 * time.Millisecond)
	balance = tmp            // write
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go addMoney(&wg)
	}

	wg.Wait()
	fmt.Println("Final Balance:", balance)
}
