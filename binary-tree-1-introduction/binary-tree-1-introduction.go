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

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(1)
	/*
		        1
		      /   \
		     NULL  NULL
	*/

	root.left = New(2)
	root.right = New(3)
	/* 2 and 3 become left and right children of 1
	           1
	         /   \
	        2      3
	     /    \    /  \
	    NULL NULL NULL NULL
	*/

	root.left.left = New(4)
	/* 4 becomes left child of 2
	           1
	       /       \
	      2          3
	    /   \       /  \
	   4    NULL  NULL  NULL
	  /  \
	NULL NULL
	*/

	fmt.Println("Root data : ", root.data)
	fmt.Println("Root->Left  data : ", root.left.data)
	fmt.Println("Root->Right data : ", root.right.data)
	fmt.Println("Root->Left->Left data : ", root.left.left.data)
}
