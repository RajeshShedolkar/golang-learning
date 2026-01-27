package main

import "fmt"

type Queue struct {
	arr      []int
	front    int
	rear     int
	size     int
	capacity int
}

func InitQueue(capacity int) *Queue {
	return &Queue{
		arr:      make([]int, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

func (q *Queue) EnQueue(data int) {
	if q.size == q.capacity {
		fmt.Println("Queue is full")
		return
	}

	q.rear = (q.rear + 1) % q.capacity
	q.arr[q.rear] = data
	q.size += 1
	fmt.Println("After EnQueue: ", q.arr)
}

func (q *Queue) DeQueue() (int, bool) {
	if q.size == 0 {
		fmt.Println("No more element left to DqQueue")
		return 0, false
	}
	deQuVal := q.arr[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size -= 1
	fmt.Println("After DeQueue: ", q.arr)
	return deQuVal, true
}

// func main(){

// 	q := InitQueue()
// 	q.EnQueue(1)
// 	q.EnQueue(2)
// 	q.EnQueue(3)
// 	q.DeQueue()
// 	q.DeQueue()
// 	q.DeQueue()
// 	q.DeQueue()
// }
