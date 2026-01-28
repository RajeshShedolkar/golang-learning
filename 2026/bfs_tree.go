package main

import "fmt"

type Node struct{
	val int
	left *Node
	right *Node
}


func NewNode(val int)*Node{
	return &Node{
		val: val,
		left: nil,
		right: nil,
	}
}

func inorder(root *Node){
	if root==nil{
		return
	}
	inorder(root.left)
	fmt.Print(root.val, " ")
	inorder(root.right)
}

func main(){
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	inorder(root)
}
