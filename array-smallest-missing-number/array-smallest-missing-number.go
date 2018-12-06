// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"
import "os"

func findFirstMissing(arr []int, start, end int) int {
	if start < 0 {
		fmt.Println("Start must be greater than 0")
		os.Exit(1)
		return 0
	}

	if start > end {
		return end + 1
	}

	if start != arr[start] {
		return start
	}

	mid := (start + end) / 2

	if arr[mid] == mid {
		return findFirstMissing(arr, mid+1, end)
	}

	return findFirstMissing(arr, start, end)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 10}
	var n int = len(arr)
	fmt.Printf("Smallest missing element is %d", findFirstMissing(arr, 0, n-1))
}
