// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

//Returns an initialized list
func (n *Node) Init() *Node {
	n.data = -1
	return n
}

//Returns an new list
func New() *Node {
	return new(Node).Init()
}

//Returns the first node in list
func (n *Node) Next() *Node {
	return n.next
}

//Returns the last node in list if exist, otherwise returns current
func (n *Node) Back() *Node {
	current := n.next
	for current != nil && current.next != nil {
		current = current.next
	}
	return current
}

//This function prints contents of linked list starting from the given node
func printList(n *Node) {
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func Push(head_ref **Node, new_data int) {
	//1. Allocate new node
	new_node := New()
	temp := *head_ref

	new_node.data = new_data
	new_node.next = *head_ref

	//2. If linked list is not NULL then set the next of last node
	if *head_ref != nil {
		for temp.next != *head_ref {
			temp = temp.next
		}
		temp.next = new_node
	} else {
		//For the fist node
		new_node.next = new_node
	}
	//4. Move the head to point to new_node
	*head_ref = new_node
}

func main() {
	//Start with the empty list
	head := New()

	Push(&head, 12)
	Push(&head, 56)
	Push(&head, 2)
	Push(&head, 11)

	fmt.Println("Created Circular LinkedList is: ")

	printList(head)
}
