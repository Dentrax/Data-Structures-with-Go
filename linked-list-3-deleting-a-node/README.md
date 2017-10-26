<h1 align="center">LinkedList Deleting a Node Source</h1>

[What It Is](#what-it-is)

## What It Is

We have discussed **[Linked List Introduction](https://github.com/Dentrax/Data-Structures-with-Go/tree/master/linked-list-1-introduction)** and **[Linked List Insertion](https://github.com/Dentrax/Data-Structures-with-Go/tree/master/linked-list-2-inserting-a-node)** in previous posts on singly linked list.
Let us formulate the problem statement to understand the deletion process. Given a `key`, delete the first occurrence of this key in linked list.

To delete a node from linked list, we need to do following steps.

* 1) Find previous node of the node to be deleted.
* 2) Changed next of previous node.
* 3) Free memory for the node to be deleted.

> * Input: Linked List = `[7 -> 1 -> 3 -> 2]`
> * Output: Created Linked List `[2 -> 3 -> 1 -> 7]`
> * Output: Linked List after Deletion of 1: `[2 -> 3 -> 7]`

> * Input: Position = 1, Linked List = `[8 -> 2 -> 3 -> 1 -> 7]`
> * Output: Linked List =  `[8 -> 3 -> 1 -> 7]`

> * Input: Position = 0, Linked List = `[8 -> 2 -> 3 -> 1 -> 7]`
> * Output: Linked List =  `[2 -> 3 -> 1 -> 7]`

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/linked-list-3-deleting-a-node/resources/deleting-a-node.png)

**Algorithm Complexity**

| Complexity		| Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n)`    |
