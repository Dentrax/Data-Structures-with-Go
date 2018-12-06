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

func TestStack(t *testing.T) {
	var testDatas = []struct {
		Stack    *Stack
		ArrayIn  []int
		ArrayOut []int
	}{
		{New(3), []int{10, 20, 30}, []int{30, 20, 10}},
		{New(5), []int{10, 20, 30, 40, 50}, []int{50, 40, 30, 20, 10}},
		{New(1), []int{7}, []int{7}},
		{New(5), []int{15, 10, 22, 33, 77}, []int{77, 33, 22, 10, 15}},
	}

	for _, data := range testDatas {
		for i := uint(0); i < data.Stack.capacity; i++ {
			Push(data.Stack, data.ArrayIn[i])
		}
		for i := uint(0); i < data.Stack.capacity; i++ {
			expected := data.ArrayOut[i]
			actual := Pop(data.Stack)

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("Stack: Expected: %d, Actual: %d", expected, actual)
			}
		}
	}
}
