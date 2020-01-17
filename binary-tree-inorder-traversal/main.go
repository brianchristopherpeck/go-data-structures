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

	inorderTraversalRecursive(root)
	fmt.Println("Inorder Traversal Recursive")

	inorderTraversalIterative(root)
	fmt.Println("Inorder Traversal Iterative")

	/*
		7 5 6 1 2 3 4 Inorder Traversal Recursive
		7 5 6 1 2 3 4 Inorder Traversal Iterative
	*/
}

func inorderTraversalRecursive(root *Tree) {

	if root != nil {

		inorderTraversalRecursive(root.Left)

		fmt.Printf("%d ",root.Value)

		inorderTraversalRecursive(root.Right)
	}

	return
}

func inorderTraversalIterative(root *Tree) {

	if root == nil {
		return
	}

	current := root
	stack := arraystack.New()
	stack.Push(current)
	for (current != nil || stack.Size() > 0) {

		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		temp, _ := stack.Pop()
		current = temp.(*Tree)

		fmt.Printf("%d ", current.Value)

		current = current.Right
	}
}