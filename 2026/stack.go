package main

import "fmt"

type Stack struct{
	arr []int
	top int
}

func InitStack() *Stack{
	return &Stack{
		arr: []int{},
		top: -1,
	}
}

func (s *Stack) Push(data int){
	s.arr = append(s.arr, data)
	s.top+=1
}

func (s *Stack) Pop()(int, bool){
	if s.top < 0{
		return 0, false
	}

	poped_val := s.arr[s.top]
	s.arr = s.arr[:len(s.arr)-1]
	s.top -= 1
	return poped_val, true
}

func test(){
	s := InitStack()
	s.Pop()
	fmt.Println(s.arr, s.top)
	s.Push(1)
	fmt.Println(s.arr, s.top)
	s.Push(2)
	fmt.Println(s.arr, s.top)
	s.Push(3)
	fmt.Println(s.arr, s.top)
	s.Pop()
	fmt.Println(s.arr, s.top)
	s.Pop()
	fmt.Println(s.arr, s.top)
	s.Pop()
	fmt.Println(s.arr, s.top)
	s.Pop()
	fmt.Println(s.arr, s.top)

}