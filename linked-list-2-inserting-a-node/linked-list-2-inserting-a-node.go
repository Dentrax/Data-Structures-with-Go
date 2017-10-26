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

func InsertAfter(prev_node *Node, new_data int) {
	//1. Check if the given prev_node is NULL
	if(prev_node == nil){
		fmt.Println("The given previous node cannot be NULL")
		return
	}

	//2. Allocate new node
	new_node := New()

	//3. Put in the data
	new_node.data = new_data

	//4. Make next of new node as next of prev_node
	new_node.next = prev_node.next

	//5. Move the next of prev_node as new_node
	prev_node.next = new_node
}

func Append(head_ref **Node, new_data int) {
	//1. Allocate new node
	new_node := New()

	last := *head_ref

	//2. Put in the data
	new_node.data = new_data

	//3. this new node is going to be last node, so make next of it as NULL
	new_node.next = nil

	//4. If the LinkedList is empty, then make the new node as head
	if(*head_ref == nil){
		*head_ref = new_node
		return
	}

	//5. Else traverse till the last node
	for(last.next != nil){
		last = last.next
	}

	//6. Change the next of last node
	last.next = new_node
	return
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
	
	//Insert 6.
	//So LinkedList becomes [NULL->6]
	Append(&head, 6)

	//Insert 7 at the beginning. 
	//So LinkedList becomes [7->NULL->6]
	Push(&head, 7)

	//Insert 1 at the beginning. 
	//So LinkedList becomes [1->7->NULL->6]
	Push(&head, 1)

	//Insert 4 at the end. 
	//So LinkedList becomes [1->7->NULL->6->4]
	Append(&head, 4)

	//Insert 8, after 7
	//So LinkedList becomes [1->7->8->NULL->6->4]
	InsertAfter(head.next, 8)

	fmt.Println("Created LinkedList is: ")

	printList(head)
}