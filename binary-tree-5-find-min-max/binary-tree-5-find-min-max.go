// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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

// Returns minimum value in a given Binary Tree
func FindMin(root *Node) int {
	//1. Check if the given node is NULL
	if root == nil {
		return MaxInt
	}

	//2. Return maximum of 3 values: 1) Root's data 2) Max in Left Subtree 3) Max in right subtree
	res := root.data
	lres := FindMin(root.left)
	rres := FindMin(root.right)

	if lres < res {
		res = lres
	}
	if rres < res {
		res = rres
	}

	return res
}

// Returns maximum value in a given Binary Tree
func FindMax(root *Node) int {
	//1. Check if the given node is NULL
	if root == nil {
		return MinInt
	}

	//2. Return maximum of 3 values: 1) Root's data 2) Max in Left Subtree 3) Max in right subtree
	res := root.data
	lres := FindMax(root.left)
	rres := FindMax(root.right)

	if lres > res {
		res = lres
	}
	if rres > res {
		res = rres
	}

	return res
}

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(2)

	root.left = New(7)
	root.right = New(5)

	root.left.right = New(6)
	root.right.right = New(9)

	root.left.right.left = New(1)
	root.left.right.right = New(11)

	root.right.right.left = New(1)

	fmt.Printf("\nMaximum element is %d", FindMax(root))
	fmt.Printf("\nMinimum element is %d", FindMin(root))
}
