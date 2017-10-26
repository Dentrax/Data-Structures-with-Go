<h1 align="center">LinkedList Inserting a Node Source</h1>

[What It Is](#what-it-is)

## What It Is

We have introduced Linked Lists in the **[previous post](https://github.com/Dentrax/Data-Structures-with-Go/tree/master/linked-list-1-introduction)**. We also created a simple linked list with 3 nodes and discussed linked list traversal.
All programs discussed in this post consider following representations of linked list .

In this post, methods to insert a new node in linked list are discussed. A node can be added in three ways ;

* 1) At the front of the linked list
* 2) After a given node.
* 3) At the end of the linked list.


Add a node at the front (A 4 steps process)
--------------------------

The new node is always added before the head of the given Linked List. And newly added node becomes the new head of the Linked List. For example if the given Linked List is `10->15->20->25` and we add an item 5 at the front, then the Linked List becomes `5->10->15->20->25`. Let us call the function that adds at the front of the list is `push()`. The `push()` must receive a pointer to the head pointer, because push must change the head pointer to point to the new node.

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/linked-list-2-inserting-a-node/resources/inserting-a-node-1.png)

**Algorithm Complexity**

| Complexity		| Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(1)`    |


Add a node after a given node (5 steps process)
--------------------------

We are given pointer to a node, and the new node is inserted after the given node.

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/linked-list-2-inserting-a-node/resources/inserting-a-node-2.png)

**Algorithm Complexity**

| Complexity		| Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(1)`    |


Add a node at the end (6 steps process)
--------------------------

The new node is always added after the last node of the given Linked List. For example if the given Linked List is `5->10->15->20->25` and we add an item 30 at the end, then the Linked List becomes `5->10->15->20->25->30`.
Since a Linked List is typically represented by the head of it, we have to traverse the list till end and then change the next of last node to new node.

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/linked-list-2-inserting-a-node/resources/inserting-a-node-3.png)

**Algorithm Complexity**

| Complexity		| Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n)`    |