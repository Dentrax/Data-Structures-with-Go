// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"

func Merge(arr []int, l, m, r int) {
	var i, j, k int
	var n1 int = m - l + 1
	var n2 int = r - m

	//Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	//Copy data to temp arrays L[] and R[]
	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	//Merge the temp arrays back into arr[l..r]
	i = 0 //Initial index of first subarray
	j = 0 //Initial index of second subarray
	k = l //Initial index of merged subarray

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	//Copy the remaining elements of L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	//Copy the remaining elements of R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func MergeSort(arr []int, l, r int) {
	if l < r {
		//Same as (l + r) / 2 but avoids overflow for large l and h
		var m int = l + (r-l)/2

		//Sort first and second halves
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		Merge(arr, l, m, r)
	}
}

func PrintArray(A []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n")
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	var arr_size int = len(arr)

	fmt.Println("Given array size is: ")
	PrintArray(arr, arr_size)

	MergeSort(arr, 0, arr_size-1)

	fmt.Println("Sorted array is: ")
	PrintArray(arr, arr_size)
}
