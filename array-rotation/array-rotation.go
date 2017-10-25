// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

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



