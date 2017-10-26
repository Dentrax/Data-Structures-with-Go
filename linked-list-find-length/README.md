<h1 align="center">LinkedList Find Length Source</h1>

[What It Is](#what-it-is)

## What It Is

For example, the function should return 5 for linked list `[1 -> 3 -> 1 -> 2 -> 1]`

![Preview Thumbnail](https://raw.githubusercontent.com/Dentrax/Data-Structures-with-Go/master/linked-list-find-length/resources/find-length.png)


METHOD 1 (Iterative Solution)
--------------------------

* 1) Initialize count as 0 
* 2) Initialize a node pointer, current = head.
* 3) Do following while current is not NULL
     a) current = current -> next
     b) count++;
* 4) Return count 

> * Input: Linked List = `[1 -> 3 -> 1 -> 2 -> 1]`
> * Output: count of nodes is `5`

**Algorithm Complexity**

| Complexity		| Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n)`    |
