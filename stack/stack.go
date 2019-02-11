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

type Stack struct {
	top      int
	capacity uint
	array    []int
}

//Returns an initialized list
func (s *Stack) Init(capacity uint) *Stack {
	s.top = -1
	s.capacity = capacity
	s.array = make([]int, capacity)
	return s
}

//Returns an new list
func New(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Stack is full when top is equal to the last index
func IsFull(stack *Stack) bool {
	return stack.top == int(stack.capacity)-1
}

// Stack is empty when top is equal to -1
func IsEmpty(stack *Stack) bool {
	return stack.top == -1
}

func Push(stack *Stack, item int) {
	if IsFull(stack) {
		return
	}
	stack.top++
	stack.array[stack.top] = item
}

func Pop(stack *Stack) int {
	if IsEmpty(stack) {
		return MinInt
	}
	temp := stack.array[stack.top]
	stack.top--
	return temp
}

func main() {
	stack := New(100)

	Push(stack, 10)
	fmt.Println("Pushed to stack : 10")
	Push(stack, 20)
	fmt.Println("Pushed to stack : 20")
	Push(stack, 30)
	fmt.Println("Pushed to stack : 30")

	fmt.Printf("Popped from stack : %d", Pop(stack))
}
