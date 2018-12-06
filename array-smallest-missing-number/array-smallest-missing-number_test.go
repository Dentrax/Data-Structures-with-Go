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

func TestFindFirstMissing(t *testing.T) {
	var testDatas = []struct {
		ArrayIn []int
		Start   int
		End     int
		Out     int
	}{
		{[]int{0, 1, 2, 4, 5}, 0, 4, 3},
		{[]int{0, 1, 2, 3, 4, 6, 8, 9, 10}, 0, 5, 5},
		{[]int{0, 1, 2, 3, 4, 6, 8, 9, 10}, 0, 8, 5},
		{[]int{0, 1, 2, 3, 4, 5, 6, 8, 9, 10}, 0, 8, 7},
		{[]int{1, 2, 3, 4, 5}, 0, 5, 0},
	}
	for _, data := range testDatas {
		expected := data.Out
		actual := findFirstMissing(data.ArrayIn, data.Start, data.End)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("FindFirstMissing: Expected: %d, Actual: %d", expected, actual)
		}
	}
}
