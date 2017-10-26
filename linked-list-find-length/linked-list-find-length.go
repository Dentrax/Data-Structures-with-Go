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

func Push(head_ref **Node, new_data int) {
	//new_node := Node{data: new_data, next: (*head_ref)}
	//*head_ref = new_node

	//1. Allocate new node
	new_node := New()

	//2. Put in the data
	new_node.data = new_data

	//3. Make next of new node as head
	new_node.next = (*head_ref)

	//4. Move the head to point to new_node
	*head_ref = new_node
}

func GetCount(head *Node) int {
	var count int = 0
	current := head

	for(current != nil){
		count++
		current = current.next
	}

	return count
}

//This function prints contents of linked list starting from the given node
func printList(n *Node){
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {
	//Start with the empty list
	head := New()
	
	Push(&head, 1)

	Push(&head, 3)

	Push(&head, 1)

	Push(&head, 2)

	Push(&head, 1)

	fmt.Printf("Count of Nodes is %d", GetCount(head))
}