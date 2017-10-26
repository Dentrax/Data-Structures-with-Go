// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"


func Swap(a *int, b *int){
	t := *a
	*a = *b
	*b = t
}

func Partition(arr []int, start, end int) int {
	
	pivot := arr[end]

	//Index of smaller element
	var i int = (start - 1)

	for j := start; j <= end - 1; j++ {
		//If current element is smaller than or equal to pivot
		if(arr[j] <= pivot){
			i++
			Swap(&arr[i], &arr[j])
		}
	}

	Swap(&arr[i + 1], &arr[end])

	return (i + 1)
}

/*The main function that implements QuickSort
	arr[] -> Array to be sorted
	start -> Starting index
	end   -> Ending index
*/
func QuickSort(arr []int, start, end int) {
	if(start < end){
		//pi is partitioning index, arr[p] is now at right place
		var pi int = Partition(arr, start, end)

		//Separately sort elements before partition and after partition
		QuickSort(arr, start, pi - 1)
		QuickSort(arr, pi + 1, end)
	}
}

func PrintArray(arr []int, size int){
	for i:=0; i < size; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	var n int = len(arr)

	QuickSort(arr, 0, n - 1)

	fmt.Println("Sorted array is: ")
	PrintArray(arr, n)
}