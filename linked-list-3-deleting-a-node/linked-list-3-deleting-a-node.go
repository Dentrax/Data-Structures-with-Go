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

func DeleteNodeWithData(head_ref **Node, delete_data int) {
	//Store head node
	temp := *head_ref
	prev := *head_ref

	//If head node itself holds the delete_data to be deleted
	if(temp != nil && temp.data == delete_data){
		*head_ref = temp.next;
		temp = nil
		return
	}

	//Search for the delete_data to be deleted, keep track of the previous node as we need to change prev->next
	for(temp != nil && temp.data != delete_data){
		prev = temp
		temp = temp.next
	}

	// If delete_data was not present in linked list
	if(temp == nil){
		return
	}

	//Unlink the node from linked list
	prev.next = temp.next

	temp = nil
}

func DeleteNodeWithPosition(head_ref **Node, delete_position int){
	//If LinkedList is empty
	if(*head_ref == nil){
		return
	}

	//Store head node
	temp := *head_ref

	//If head needs to be removed
	if(delete_position == 0){
		*head_ref = temp.next
		temp = nil
		return
	}

	//Find previous node of the node to be deleted
	for i := 0; temp != nil && i < delete_position - 1; i++ {
		temp = temp.next
	}

	//If position is more than number of nodes
	if(temp == nil || temp.next == nil){
		return
	}

	//Node temp->next is the node to be deleted, Store pointer to the next of node to be deleted
	next := temp.next.next

	//Unlink the node from linked list
	temp.next = nil

	//Unlink the deleted node from list
	temp.next = next
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
	
	Push(&head, 7)

	Push(&head, 1)

	Push(&head, 3)

	Push(&head, 2)

	Push(&head, 8)

	fmt.Println("Created LinkedList is: ")

	printList(head)

	DeleteNodeWithData(&head, 1)

	DeleteNodeWithPosition(&head, 4)

	fmt.Println("LinkedList after Deletion of 1: ")

	printList(head)
}