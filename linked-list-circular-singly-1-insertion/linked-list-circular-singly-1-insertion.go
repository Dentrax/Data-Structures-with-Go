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

//Function to insert node in an empty List
func AddToEmpty(last *Node, data int) *Node {
	if(last != nil){
		return last
	}

	//Creating a node
	temp := New()

	//Assigning the data
	temp.data = data
	last = temp

	//Creating the link
	last.next = last

	return last
}

//Function to insert node in the beginning of the List
func AddBegin(last *Node, data int) *Node {
	if(last == nil){
		return AddToEmpty(last, data)
	}

	//Creating a node
	temp := New()

	//Assigning the data
	temp.data = data

	//List was empty. We link single node to itself
	temp.next = last.next
	last.next = temp

	return last
}

//Function to insert node in the end of the List
func AddEnd(last *Node, data int) *Node {
	if(last == nil){
		return AddToEmpty(last, data)
	}

	//Creating a node
	temp := New()

	//Assigning the data
	temp.data = data

	//List was empty. We link single node to itself
	temp.next = last.next
	last.next = temp
	last = temp

	return last
}

//Function to insert node in the end of the List
func AddAfter(last *Node, data int, item int) *Node {
	if(last == nil){
		return nil
	}

	//Creating a node
	temp := New()
	p := New()

	p = last.next

	for(p != last.next){
		if(p.data == item){
			//Assigning the data
			temp.data = data
			//Adjusting the links
			temp.next = p.next

			//Adding newly allocated node after p
			p.next = temp

			//Checking for the last node
			if(p == last){
				last = temp
			}

			return last
		}
		p = p.next
	}

	fmt.Printf("\n%d not present in the list.\n", item)

	return last
}

func Traverse(last *Node){
	p := New()

	//If list is empty, return
	if(last == nil){
		fmt.Println("List is empty")
		return
	}

	//Pointing to first Node of the list
	p = last.next

	//Traversing the list
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

//This function prints contents of linked list starting from the given node
func printList(n *Node){
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {
	last := New()

	last = AddEnd(last, 12)
	last = AddToEmpty(last, 6)
	last = AddBegin(last, 4)
	last = AddBegin(last, 2)
	last = AddEnd(last, 8)
	last = AddAfter(last, 10, 8)

	Traverse(last)
}