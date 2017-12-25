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
func DeleteTree(node *Node) {
	//1. Check if the given node is NULL
	if node == nil {
		return
	}

	//2. First delete both subtrees
	DeleteTree(node.left)
	DeleteTree(node.right)

	//3. Then delete the node
	fmt.Printf("\nDeleting node: %d", node.data)
	//free(node);
	node = nil
}

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(1)

	root.left = New(2)
	root.right = New(3)

	root.left.left = New(4)
	root.left.right = New(5)

	DeleteTree(root)
	root = nil

	fmt.Println("\nTree deleted")
}
