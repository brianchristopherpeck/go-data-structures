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
		  5	   3
		 /\	  / \
		7 6	 2   4

	*/

	root := &Tree{ Value: 1,}
	root.Right = &Tree{ Value: 3,}
	root.Left = &Tree{ Value: 5,}
	root.Left.Left = &Tree{ Value: 7,}
	root.Left.Right = &Tree{ Value: 6,}
	root.Right.Left = &Tree{ Value: 2,}
	root.Right.Right = &Tree{ Value: 4,}

	preorderTraversalRecursive(root)
	fmt.Println("Preorder Traversal Recursive")

	preorderTraversalIterative(root)
	fmt.Println("Preorder Traversal Iterative")

	/*
		1 5 7 6 3 2 4 Preorder Traversal Recursive
		1 5 7 6 3 2 4 Preorder Traversal Iterative
	*/
}

func preorderTraversalRecursive(root *Tree) {

	if root != nil {
		fmt.Printf("%d ",root.Value)

		preorderTraversalRecursive(root.Left)

		preorderTraversalRecursive(root.Right)
	}

	return
}

func preorderTraversalIterative(root *Tree) {

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

		if current.Right != nil {
			stack.Push(current.Right)
		}

		if current.Left != nil {
			stack.Push(current.Left)
		}
	}

	for i := 0; i < len(output); i++ {
		fmt.Printf("%d " , output[i])
	}
}