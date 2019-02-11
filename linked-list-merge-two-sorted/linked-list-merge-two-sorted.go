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

//Pull off the front node of the source and put it in dest
/* MoveNode() function takes the node from the front of the
   source, and move it to the front of the dest.
   It is an error to call this with the source list empty.

   Before calling MoveNode():
   source == {1, 2, 3}
   dest == {1, 2, 3}

   Affter calling MoveNode():
   source == {2, 3}
   dest == {1, 1, 2, 3} */
func MoveNode(dest_ref **Node, source_ref **Node) {
	//The front source code
	new_node := *source_ref

	//Advance the source pointer
	*source_ref = new_node.next

	//Link the old dest off the new node
	new_node.next = *dest_ref

	//Move dest to point to the new node
	*dest_ref = new_node
}

/* Takes two lists sorted in increasing order, and splices
   their nodes together to make one big sorted list which
   is returned.  */
func SortedMerge(a *Node, b *Node) *Node {
	//A dummy first node to hang the result on
	dummy := New()

	//Tail points to the last result node
	tail := dummy

	//So tail.next is the place to add new nodes to result
	dummy.next = nil

	for {
		if a == nil {
			//If either list runs out, use the other list
			tail.next = b
			break
		} else if b == nil {
			tail.next = a
			break
		}

		if a.data <= b.data {
			MoveNode(&(tail.next), &a)
		} else {
			MoveNode(&(tail.next), &b)
		}

		tail = tail.next
	}

	return dummy.next
}

//This function prints contents of linked list starting from the given node
func printList(n *Node) {
	for n != nil {
		fmt.Print(n.data)
		fmt.Print(" ,")
		n = n.next
	}
}

func GetDataList(n *Node) []int {
	data := []int{}
	for n != nil {
		data = append(data, n.data)
		n = n.next
	}
	return data
}

func main() {
	//Start with the empty list
	res := New()
	a := New()
	b := New()

	Push(&a, 15)
	Push(&a, 10)
	Push(&a, 5)

	Push(&b, 20)
	Push(&b, 3)
	Push(&b, 2)

	//Remove duplicates from LinkedList
	res = SortedMerge(a, b)

	fmt.Println("Merged LinkedList is:")
	printList(res)
}
