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

func TestBinaryTreeNew(t *testing.T) {
	var testDatas = []struct {
		Node  *Node
		Data  int
		Left  *Node
		Right *Node
	}{
		{New(-1), -1, New(-1), New(-1)},
		{New(0), 0, New(0), New(0)},
		{New(1), 1, New(1), New(1)},
		{New(7), 7, New(7), New(7)},
	}

	for _, data := range testDatas {
		data.Node.right = data.Right
		data.Node.left = data.Left

		expected := data.Data

		actualLeft := data.Node.left.data
		actualRight := data.Node.right.data

		if !reflect.DeepEqual(expected, actualLeft) || !reflect.DeepEqual(expected, actualRight) {
			t.Errorf("BinaryTreeNew: Expected: %d, ActualLeft: %d, ActualRight: %d", expected, actualLeft, actualRight)
		}
	}
}
