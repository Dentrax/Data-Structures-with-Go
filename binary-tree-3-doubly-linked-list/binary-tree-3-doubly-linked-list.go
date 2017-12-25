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

// Standard Inorder traversal of tree
func InOrder(root *Node) {
	//1. Check if the given node is NULL
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d ", root.data)
	InOrder(root.right)
}

// Changes left pointers to work as previous pointers in converted DLL
// The function simply does inorder traversal of Binary Tree and updates
// left pointer using previously visited node
var pre *Node = nil

func FixPrevPtr(root *Node) {
	//1. Check if the given root is NULL
	if root == nil {
		return
	}

	FixPrevPtr(root.left)
	root.left = pre
	pre = root
	FixPrevPtr(root.right)
}

// Changes right pointers to work as next pointers in converted DLL
var prev *Node = nil

func FixNextPtr(root *Node) *Node {
	// Find the right most node in BT or last node in DLL
	for root != nil && root.right != nil {
		root = root.right
	}

	// Start from the rightmost node, traverse back using left pointers.
	// While traversing, change right pointer of nodes.
	for root != nil && root.left != nil {
		prev = root
		root = root.left
		root.right = prev
	}

	// The leftmost node is head of linked list, return it
	return root
}

func BTToDLL(root *Node) *Node {
	// Set the previous pointer
	FixPrevPtr(root)

	// Set the next pointer and return head of DLL
	return FixNextPtr(root)
}

func PrintList(root *Node) {
	for root != nil {
		fmt.Printf("%d ", root.data)
		root = root.right
	}
}

func main() {

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(10)

	root.left = New(12)
	root.right = New(15)

	root.left.left = New(25)
	root.left.right = New(30)

	root.right.left = New(36)

	fmt.Println("\nInorder Tree Traversal")
	InOrder(root)

	head := BTToDLL(root)
	fmt.Println("\nDLL Traversal")
	PrintList(head)
}
