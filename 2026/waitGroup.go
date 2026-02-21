package main

import (
	"fmt"
	"sync"
)

type MyWaitGroup struct {
	counter int
	mutex   sync.Mutex
	cond    *sync.Cond
}

func NewMyWaitGroup() *MyWaitGroup {
	wg := &MyWaitGroup{}
	wg.cond = sync.NewCond(&wg.mutex)
	return wg
}

func (wg *MyWaitGroup) Add(n int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.counter += n
	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.counter == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *MyWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *MyWaitGroup) Wait() {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	for wg.counter > 0 {
		wg.cond.Wait()
	}
}

func worker(id int, wg *MyWaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id, "started")
	// simulate work
}

func main() {
	wg := NewMyWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
