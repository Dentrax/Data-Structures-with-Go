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

//Function to search a given key in a given BST
func Search(root *Node, key int) *Node {
	//1. Base Cases: root is null or key is present at root
	if root == nil || root.data == key {
		fmt.Println("The given previous node cannot be NULL")
		return root
	}

	//2. Key is greater than root's key
	if root.data < key {
		return Search(root.right, key)
	}

	//3. Key is smaller than root's key
	return Search(root.left, key)
}

//A utility function to do inorder traversal of BST
func PrintInOrder(root *Node) {
	if root != nil {
		PrintInOrder(root.left)
		fmt.Printf("%d \n", root.data)
		PrintInOrder(root.right)
	}
}

//A utility function to insert a new node with given key in BST
func Insert(node *Node, key int) *Node {
	//1. If the tree is empty, return a new node
	if node == nil {
		return New(key)
	}

	//2. Otherwise, recur down the tree
	if key < node.data {
		node.left = Insert(node.left, key)
	} else if key > node.data {
		node.right = Insert(node.right, key)
	}

	//3. Return the (unchanged) node pointer
	return node
}

func main() {
	/* Let us create following BST
	              50
	           /     \
	          30      70
	         /  \    /  \
	       20   40  60   80 */

	//To allocate dynamically a new Node in C language : root = (struct Node*) malloc(sizeof(struct Node));
	root := New(50)
	Insert(root, 30)
	Insert(root, 20)
	Insert(root, 40)
	Insert(root, 70)
	Insert(root, 60)
	Insert(root, 80)

	//Print inoder traversal of the BST
	PrintInOrder(root)
}
