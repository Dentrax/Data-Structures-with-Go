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

func TestLeftRotate(t *testing.T) {
	var testDatas = []struct {
		ArrayIn  []int
		Count    int
		Depth    int
		ArrayOut []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 7, []int{3, 4, 5, 6, 7, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 6, []int{3, 4, 5, 6, 1, 2, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 2, []int{2, 1, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6, []int{2, 3, 4, 5, 6, 1, 7}},
	}
	for _, data := range testDatas {
		expected := data.ArrayOut
		leftRotate(data.ArrayIn, data.Count, data.Depth)
		actual := data.ArrayIn

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("LeftRotate: Expected: %d, Actual: %d", expected, actual)
		}
	}
}
