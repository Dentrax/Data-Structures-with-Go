// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import "fmt"

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func minDist (arr []int, n, x, y int) int {
	var min_dist int = MaxInt

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if((x == arr[i] && y == arr[j] || y == arr[i] && x == arr[j]) && min_dist > Abs(i - j)) {
				min_dist = Abs(i - j)
			}
		}
	}

	return min_dist
}

func main() {
	arr := []int{3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}
	var n int = len(arr)
	var x int = 3
	var y int = 6
	fmt.Printf("Minimum distance between %d and %d is %d", x, y, minDist(arr, n, x, y))
}