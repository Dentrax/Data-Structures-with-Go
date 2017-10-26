<h1 align="center">Merge Two Sorted Source</h1>

[What It Is](#what-it-is)

## What It Is

Write a `SortedMerge()` function that takes two lists, each of which is sorted in increasing order, and merges the two together into one list which is in increasing order. `SortedMerge()` should return the new list. The new list should be made by splicing
together the nodes of the first two lists.

> * Input: LinkedList a `5->10->15`
> * Input: LinkedList b `2->3->20`
> * Output: `2->3->5->10->15->20`

There are many cases to deal with: either `a` or `b` may be empty, during processing either `a` or `b` may run out first, and finally there’s the problem of starting the result list empty, and building it up while going through `a` and `b`.

METHOD 1 (Using Dummy Nodes)
--------------------------

The strategy here uses a temporary dummy node as the start of the result list. The pointer Tail always points to the last node in the result list, so appending new nodes is easy.
The dummy node gives tail something to point to initially when the result list is empty. This dummy node is efficient, since it is only temporary, and it is allocated in the stack. The loop proceeds, removing one node from either `a` or `b`, and adding it to tail. When
we are done, the result is in dummy.next

**Algorithm Complexity**

| Complexity	    | Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n)`    |