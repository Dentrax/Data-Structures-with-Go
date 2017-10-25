<h1 align="center">Array Rotation Source </h1>

[What It Is](#what-it-is)

[How To Use](#how-to-use)

## What It Is

-Write a function rotate(ar[], d, n) that rotates `arr[]` of size `n` by `d` elements
-1 2 3 4 5 6 7

-Rotation of the above array by 2 will make array
-3 4 5 6 7 1 2

**METHOD 1 (Use temp array)**

Input arr[] = [1, 2, 3, 4, 5, 6, 7], d = 2, n =7
1) Store `d` elements in a temp array
   temp[] = [1, 2]
2) Shift rest of the `arr[]`
   arr[] = [3, 4, 5, 6, 7, 6, 7]
3) Store back the `d` elements
   arr[] = [3, 4, 5, 6, 7, 1, 2]

Algorithm Complexity
--------------------------

| Complexity		    | Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n)`    |
| `Auxiliary Space` | `O(d)`    |


**METHOD 2 (Rotate one by one)**

```go
leftRotate(arr[], d, n)
start
  For i = 0 to i < d
    Left rotate all elements of arr[] by one
end
```

To rotate by one, store arr[0] in a temporary variable temp, move arr[1] to arr[0], arr[2] to arr[1] …and finally temp to arr[n-1]
Let us take the same example arr[] = [1, 2, 3, 4, 5, 6, 7], d = 2
Rotate arr[] by one 2 times
We get `[2, 3, 4, 5, 6, 7, 1]` after first rotation and `[3, 4, 5, 6, 7, 1, 2]` after second rotation.

Algorithm Complexity
--------------------------

| Complexity		    | Notation  |
| ----------------- |:---------:|
| `Time Complexity`	| `O(n*d)`  |
| `Auxiliary Space` | `O(1)`    |


## How To Use

