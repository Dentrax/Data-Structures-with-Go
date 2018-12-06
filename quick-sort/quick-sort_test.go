// ====================================================
// Data-Structures-with-Go Copyright(C) 2017 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var testDatas = []struct {
		ArrayIn  []int
		ArrayOut []int
	}{
		{[]int{1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{3, 2, 1, 4}, []int{1, 2, 3, 4}},
		{[]int{9, 8, 6, 5, 7, 4, 3, 0, 2, 1}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{-3, -2, -1, -4, 0}, []int{-4, -3, -2, -1, 0}},
	}
	for _, data := range testDatas {
		expected := data.ArrayOut
		QuickSort(data.ArrayIn, 0, len(data.ArrayIn)-1)
		actual := data.ArrayIn

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("QuickSort: Expected: %d, Actual: %d", expected, actual)
		}
	}
}
