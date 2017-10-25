// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

/*** INFO

-Write a function rotate(ar[], d, n) that rotates arr[] of size n by d elements
-1 2 3 4 5 6 7

-Rotation of the above array by 2 will make array
-3 4 5 6 7 1 2

-METHOD 1 (Use temp array) :

Input arr[] = [1, 2, 3, 4, 5, 6, 7], d = 2, n =7
1) Store d elements in a temp array
   temp[] = [1, 2]
2) Shift rest of the arr[]
   arr[] = [3, 4, 5, 6, 7, 6, 7]
3) Store back the d elements
   arr[] = [3, 4, 5, 6, 7, 1, 2]

-Time complexity O(n)
-Auxiliary Space: O(d)


-METHOD 2 (Rotate one by one) :

leftRotate(arr[], d, n)
start
  For i = 0 to i < d
    Left rotate all elements of arr[] by one
end

To rotate by one, store arr[0] in a temporary variable temp, move arr[1] to arr[0], arr[2] to arr[1] …and finally temp to arr[n-1]
Let us take the same example arr[] = [1, 2, 3, 4, 5, 6, 7], d = 2
Rotate arr[] by one 2 times
We get [2, 3, 4, 5, 6, 7, 1] after first rotation and [ 3, 4, 5, 6, 7, 1, 2] after second rotation.

-Time complexity: O(n*d)
-Auxiliary Space: O(1)

***/

package main

import "fmt"

func leftRotate(arr []int, d int, n int){
	for i := 0; i < d; i++ {
		leftRotateByOne(arr, n)
	}
}

func leftRotateByOne(arr []int, n int){
	var i, temp int
	temp = arr[0]
	for i = 0; i < n - 1; i++ {
		arr[i] = arr[i + 1]
	}
	arr[i] = temp
}

func printArray(arr []int, size int){
	if(len(arr) < size){
		fmt.Println("[ArrayRotation::printArray()] -> Index out of range. Max : ", len(arr))
		return;
	}
	for i := 0; i < size; i++ {
		fmt.Print(arr[i])
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("-INPUT-")
	printArray(arr, len(arr) - 1)
	leftRotate(arr, 2 , 7)
	fmt.Println()
	fmt.Println("-OUTPUT-")
	printArray(arr, 7)
}



