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

func TestLinkedListReverse(t *testing.T) {
	var testDatas = []struct {
		Node     *Node
		ArrayIn  []int
		ArrayOut []int
	}{
		{New(), []int{20, 4, 15, 85}, []int{-1, 20, 4, 15, 85, -1}},
		{New(), []int{10, 20, 30, 40, 50}, []int{-1, 10, 20, 30, 40, 50, -1}},
		{New(), []int{7}, []int{-1, 7, -1}},
	}

	for _, data := range testDatas {
		for i := 0; i < len(data.ArrayIn); i++ {
			Push(&data.Node, data.ArrayIn[i])
		}

		Reverse(&data.Node)

		n := data.Node
		i := 0
		for n != nil {
			expected := data.ArrayOut[i]
			actual := n.data

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("LinkedListReverse: Expected: %d, Actual: %d", expected, actual)
			}

			n = n.Next()
			i++
		}
	}
}
