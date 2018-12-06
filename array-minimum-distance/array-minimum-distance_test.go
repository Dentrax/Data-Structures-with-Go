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

func TestAbs(t *testing.T) {
	var testDatas = []struct {
		Input  int
		Output int
	}{
		{1, 1},
		{7, 7},
		{-1, 1},
		{-7, 7},
		{0, 0},
	}
	for _, data := range testDatas {
		expected := data.Output
		actual := Abs(data.Input)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Abs: Expected: %d, Actual: %d", expected, actual)
		}

	}
}

func TestMinDist(t *testing.T) {
	var testDatas = []struct {
		Array    []int
		Num1     int
		Num2     int
		Distance int
	}{
		{[]int{3, 4, 5, 6}, 3, 6, 3},
		{[]int{1, 2, 3}, 1, 3, 2},
		{[]int{9, 7, 5, 3, 1}, 1, 7, 3},
		{[]int{-1, -7, 5, 5, 1, 5}, -1, 5, 2},
		{[]int{-1, -7, 5, 5, -7, -1}, -1, -7, 1},
	}
	for _, data := range testDatas {
		expected := data.Distance
		actual := minDist(data.Array, len(data.Array), data.Num1, data.Num2)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("MinDist: Expected: %d, Actual: %d", expected, actual)
		}
	}
}
