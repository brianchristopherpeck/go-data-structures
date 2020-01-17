package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type Tree struct {
	Value int
	Left *Tree
	Right *Tree
}

func main() {
	/*
			 1
			/ \
		  5	   2
		 /\	  / \
		7 6	 3   4

	*/

	root := &Tree{ Value: 1,}
	root.Right = &Tree{ Value: 3,}
	root.Left = &Tree{ Value: 5,}
	root.Left.Left = &Tree{ Value: 7,}
	root.Left.Right = &Tree{ Value: 6,}
	root.Right.Left = &Tree{ Value: 2,}
	root.Right.Right = &Tree{ Value: 4,}

	postorderTraversalRecursive(root)
	fmt.Println("Postorder Traversal Recursive")

	postorderTraversalIterative(root)
	fmt.Println("Postorder Traversal Iterative")

	/*
		7 6 5 2 4 3 1 Postorder Traversal Recursive
		7 6 5 2 4 3 1 Postorder Traversal Iterative
	*/
}

func postorderTraversalRecursive(root *Tree) {

	if root != nil {

		postorderTraversalRecursive(root.Left)

		postorderTraversalRecursive(root.Right)

		fmt.Printf("%d ",root.Value)
	}

	return
}

func postorderTraversalIterative(root *Tree) {

	if root == nil {
		return
	}

	output := []int{}

	current := root
	stack := arraystack.New()
	stack.Push(current)
	for stack.Size() > 0 {
		temp, _ := stack.Pop()
		current = temp.(*Tree)

		output = append(output, current.Value)
		//fmt.Printf("%d ", current.Value)

		if current.Left != nil {
			stack.Push(current.Left)
		}

		if current.Right != nil {
			stack.Push(current.Right)
		}
	}

	for i := len(output) - 1; i >= 0; i-- {
		fmt.Print(output[i])
		fmt.Print(" ")
	}
}