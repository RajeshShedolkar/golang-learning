package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()         // 🔒 lock
	counter++
	mu.Unlock()       // 🔓 unlock
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// wg.Wait()
	defer fmt.Println("Final Counter:", counter)
}
