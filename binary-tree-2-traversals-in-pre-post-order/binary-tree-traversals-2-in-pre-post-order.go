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

func PrintPostOrder(node *Node) {
	//1. Check if the given node is NULL
	if node == nil {
		return
	}

	//2. First recur on left subtree
	PrintPostOrder(node.left)

	//3. Then recur on right subtree
	PrintPostOrder(node.right)

	//4. Now deal with the node
	fmt.Printf("%d ", node.data)
}

func PrintInOrder(node *Node) {
	//1. Check if the given node is NULL
	if node == nil {
		return
	}

	//2. First recur on left child
	PrintInOrder(node.left)

	//3. Then print the data of node
	fmt.Printf("%d ", node.data)

	//4. Now recur on right child
	PrintInOrder(node.right)
}

func PrintPreOrder(node *Node) {
	//1. Check if the given node is NULL
	if node == nil {
		return
	}

	//2. First print data of node
	fmt.Printf("%d ", node.data)

	//3. Then recur on left sutree
	PrintPreOrder(node.left)

	//4. Now recur on right subtree
	PrintPreOrder(node.right)
}

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(1)

	root.left = New(2)
	root.right = New(3)

	root.left.left = New(4)
	root.left.right = New(5)

	fmt.Println("\nPreorder traversal of binary tree is :")
	PrintPreOrder(root)
	fmt.Println("\nInorder traversal of binary tree is :")
	PrintInOrder(root)
	fmt.Println("\nPostorder traversal of binary tree is :")
	PrintPostOrder(root)
}
