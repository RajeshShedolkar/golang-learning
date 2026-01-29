package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

// type Queue struct{
// 	arr []int
// 	capacity int
// 	size int
// 	front int
// 	rear int
// }

// func (q *Queue) EnQueue(){

// }

// func (q *Queue) DeQueue(){

// }

func NewNode(val int) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Print(root.val, " ")
	inorder(root.right)
}

func LevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Print(curr.val, " ")

		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}

func main() {
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(12)
	root.left.left.left = NewNode(5)
	//inorder(root)
	LevelOrder(root)
}
