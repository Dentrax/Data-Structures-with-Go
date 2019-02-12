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

func TestMergeTwoSortedNew(t *testing.T) {
	var testDatas = []struct {
		ArrayIn1 []int
		ArrayIn2 []int
		Out      []int
	}{
		{[]int{20, 4, 15, 85}, []int{6, 7}, []int{20, 4, 15, 85, -1, 6, 7}},
		{[]int{10, 20, 30}, []int{5, 6}, []int{10, 20, 30, -1, 5, 6}},
		{[]int{7}, []int{7}, []int{7, -1, 7}},
		{[]int{}, []int{}, []int{-1}},
	}

	for _, data := range testDatas {

		a := New()
		b := New()

		for i := 0; i < len(data.ArrayIn1); i++ {
			Push(&a, data.ArrayIn1[i])
		}

		for i := 0; i < len(data.ArrayIn2); i++ {
			Push(&b, data.ArrayIn2[i])
		}

		actual := SortedMerge(a, b)

		expected := New()

		for i := 0; i < len(data.Out); i++ {
			Push(&expected, data.Out[i])
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("MergeTwoSorted: Expected: %v, Actual: %v", GetDataList(expected), GetDataList(actual))
		}
	}
}
