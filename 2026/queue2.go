package main

import (
	"fmt"
)

type Q struct {
	arr      []int
	capacity int
	size     int
	front    int
	rear     int
}

func NewQ(capacity int) *Q {
	if capacity > 0 {
		return &Q{
			arr:      make([]int, capacity),
			capacity: capacity,
			size:     0,
			front:    0,
			rear:     -1,
		}
	}
	return nil
}

func (q *Q) EnQueue(val int) error {
	if q.size == q.capacity {
		fmt.Println("No EnQueue possible, Queue is full")
		return fmt.Errorf("queue full")
	}
	q.rear = (q.rear + 1) % q.capacity
	q.arr[q.rear] = val
	q.size += 1
	return nil
}

func (q *Q) DeQueue() (int, error) {
	if q.size == 0 {
		fmt.Println("Queue is empty, No more DeQueue possible")
		return 0, fmt.Errorf("Queue is Empty")
	}
	popped_val := q.arr[q.front]
	q.arr[q.front] = 0
	q.front = (q.front + 1) % q.capacity
	q.size -= 1
	return popped_val, nil
}

func main() {
	fmt.Println("---- Create Queue ----")
	q := NewQ(5)
	fmt.Println("Initial:", q.arr, "size:", q.size)

	fmt.Println("\n---- Dequeue on empty ----")
	_, err := q.DeQueue()
	fmt.Println("Error:", err)

	fmt.Println("\n---- Enqueue till full ----")
	for i := 1; i <= 5; i++ {
		err := q.EnQueue(i * 10)
		fmt.Println("Enqueue", i*10, "Error:", err)
		fmt.Println("Queue:", q.arr, "front:", q.front, "rear:", q.rear, "size:", q.size)
	}

	fmt.Println("\n---- Enqueue when full ----")
	err = q.EnQueue(999)
	fmt.Println("Error:", err)

	fmt.Println("\n---- Dequeue 2 elements ----")
	for i := 0; i < 2; i++ {
		val, err := q.DeQueue()
		fmt.Println("Dequeued:", val, "Error:", err)
		fmt.Println("Queue:", q.arr, "front:", q.front, "rear:", q.rear, "size:", q.size)
	}

	fmt.Println("\n---- Enqueue again (circular behavior) ----")
	q.EnQueue(60)
	q.EnQueue(70)
	fmt.Println("Queue:", q.arr, "front:", q.front, "rear:", q.rear, "size:", q.size)

	fmt.Println("\n---- Dequeue all ----")
	for q.size > 0 {
		val, _ := q.DeQueue()
		fmt.Println("Dequeued:", val)
		fmt.Println("Queue:", q.arr, "front:", q.front, "rear:", q.rear, "size:", q.size)
	}

	fmt.Println("\n---- Dequeue again on empty ----")
	_, err = q.DeQueue()
	fmt.Println("Error:", err)
}
