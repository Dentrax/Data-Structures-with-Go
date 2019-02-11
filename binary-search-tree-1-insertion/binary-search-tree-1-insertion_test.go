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

func TestBinarySearchTreeNew(t *testing.T) {
	var testDatas = []struct {
		Node    *Node
		ArrayIn []int
		Test    int
		Out     int
	}{
		{New(0), []int{20, 4, 15, 85}, 15, 15},
		{New(3), []int{10, 20, 30, 40, 50}, 50, 50},
		{New(7), []int{7}, 7, 7},
	}

	for _, data := range testDatas {
		for i := 0; i < len(data.ArrayIn); i++ {
			Insert(data.Node, data.ArrayIn[i])
		}

		actual := Search(data.Node, data.Test)

		expected := data.Out

		if !reflect.DeepEqual(expected, actual.data) {
			t.Errorf("BinarySearchTree: Expected: %d, Actual: %d", expected, actual)
		}
	}
}
