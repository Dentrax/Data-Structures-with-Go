// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

//Returns an initialized list
func (n *Node) Init(data int) *Node {
	n.data = data
	n.left = nil
	n.right = nil
	return n
}

//Returns an new list
func New(data int) *Node {
	return new(Node).Init(data)
}

/* Compute the "height" of a tree -- the number of
    nodes along the longest path from the root node
    down to the farthest leaf node.*/
func GetHeight(node *Node) int {
	//1. Check if the given node is NULL
	if node == nil {
		return 0
	}

	//2. Compute the height of each subtree
	lheight := GetHeight(node.left)
	rheight := GetHeight(node.right)

	//3. Use the larger one
	if lheight > rheight {
		return lheight + 1
	}
	return rheight + 1
}

/* Print nodes at a given level */
func PrintGivenLevel(root *Node, level int) {
	//1. Check if the given root is NULL
	if root == nil {
		return
	}

	if level == 1 {
		fmt.Printf("%d ", root.data)
	} else if level > 1 {
		PrintGivenLevel(root.left, level-1)
		PrintGivenLevel(root.right, level-1)
	}
}

/* Function to print level order traversal a tree*/
func PrintLevelOrder(root *Node) {
	//1. Check if the given root is NULL
	if root == nil {
		fmt.Println("The given root node cannot be NULL")
		return
	}

	h := GetHeight(root)

	for i := 1; i <= h; i++ {
		PrintGivenLevel(root, i)
	}
}

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(1)

	root.left = New(2)
	root.right = New(3)

	root.left.left = New(4)
	root.left.right = New(5)

	fmt.Println("\nLevel Order traversal of binary tree is :")
	PrintLevelOrder(root)
}
