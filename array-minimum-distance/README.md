<h1 align="center">Array Minimum Distance Source</h1>

[What It Is](#what-it-is)

## What It Is

* Given an unsorted array `arr[]` and two numbers x and y, find the minimum distance between `x` and `y` in `arr[]`. The array might also contain duplicates. You may assume that both `x` and `y` are different and present in `arr[]`.

Examples
--------------------------

> * Input: arr[] = {1, 2}, x = 1, y = 2
> * Output: Minimum distance between 1 and 2 is 1.

> * Input: arr[] = {3, 4, 5}, x = 3, y = 5
> * Output: Minimum distance between 3 and 5 is 2.

> * Input: arr[] = {3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}, x = 3, y = 6
> * Output: Minimum distance between 3 and 6 is 4.

> * Input: arr[] = {2, 5, 3, 5, 4, 4, 2, 3}, x = 3, y = 2
> * Output: Minimum distance between 3 and 2 is 1.

METHOD 1 (Simple)
--------------------------

Use two loops: The outer loop picks all the elements of arr[] one by one. The inner loop picks all the elements after the element picked by outer loop. If the elements picked by outer and inner loops have same values as x or y then if needed update the minimum distance calculated so far.

**Algorithm Complexity**

| Complexity	    | Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n^2)`  |