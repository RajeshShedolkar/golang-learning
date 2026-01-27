package main
import "fmt"

type Queue struct{
	arr []int
	front int
	rear int
	size int
	capacity int
}

func InitQueue(capacity int) *Queue{
	return &Queue{
		arr: make([]int, capacity),
		front: 0,
		rear: -1,
		size: 0,
		capacity: capacity,
	}
}

func (q *Queue) EnQueue(data int){
	q.arr = append(q.arr, data)
	fmt.Println("After EnQueue: ", q.arr)
}

func (q *Queue) DeQueue()(int, bool){
	if len(q.arr)==0{
		fmt.Println("No more element left to DqQueue")
		return 0, false
	}
	deQuVal := q.arr[0]
	q.arr = q.arr[1:]
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