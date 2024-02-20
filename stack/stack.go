package stack // Package for stack implementation

import (
	"errors"
	"fmt"
)

// node represents a single element in a linked list stack
type node[T int | float32 | float64] struct {
	Next *node[T] // Pointer to the next node
	Val  T        // Value stored in the node
}

const StackMaxSize = 100 // Maximum size for array-based stack

// stack interface defines common operations for stack implementations
type stack[T int | float32 | float64] interface {
	Pop() (T, error) // Removes and returns the top element
	Push(T) error    // Adds an element to the top of the stack
	Top() int        // Returns the index of the top element
	Print()          // Prints the contents of the stack
	IsEmpty() bool   // Returns true if length is 0
}

// stackArray implements a stack using a fixed-size array
type stackArray[T int | float32 | float64] struct {
	arr [StackMaxSize]T // Array to hold stack elements
	top int             // Index of the top element
}

// Push adds an element to the top of the array stack
func (stack *stackArray[T]) Push(element T) error {
	if stack.top == StackMaxSize-1 {
		return errors.New("Stack Overflow!!!")
	}

	stack.arr[stack.top+1] = element
	stack.top++
	return nil
}

// Pop removes and returns the top element from the array stack
func (stack *stackArray[T]) Pop() (T, error) {
	if stack.top == -1 {
		return 0, errors.New("Stack Underflow!!!")
	}

	stack.top--
	return stack.arr[stack.top+1], nil
}

// Print prints the contents of the array stack
func (stack *stackArray[T]) Print() {
	fmt.Print("[ ")
	for i := stack.top; i >= 0; i-- { // Iterate from top to bottom
		fmt.Print(stack.arr[i], " ")
	}
	fmt.Print("]\n")
}

// NewArray creates a new instance of an array stack
func NewArray[T int | float32 | float64]() stackArray[T] {
	return stackArray[T]{top: -1}
}

// Top returns the index of the top element in the array stack
func (stack *stackArray[T]) Top() int {
	return stack.top
}

// stackList implements a stack using a linked list
type stackList[T int | float32 | float64] struct {
	topNode  *node[T] // Pointer to the top node
	top      int      // Index of the top element
	capacity int      // Maximum capacity of the stack
}

// Push adds an element to the top of the linked list stack
func (stack *stackList[T]) Push(element T) error {
	if stack.top == stack.capacity-1 {
		return errors.New("Stack Overflow!!!")
	}

	// Create a new node and make it the new top node
	newTop := new(node[T])
	newTop.Val = element
	newTop.Next = stack.topNode
	stack.topNode = newTop

	stack.top++
	return nil
}

// Pop removes and returns the top element from the linked list stack
func (stack *stackList[T]) Pop() (T, error) {
	if stack.top == -1 {
		return 0, errors.New("Stack Underflow!!!")
	}

	// Remove the top node and return its value
	store := stack.topNode
	stack.topNode = stack.topNode.Next
	stack.top--

	return store.Val, nil
}

// NewList creates a new instance of a linked list stack
func NewList[T int | float32 | float64](capacity ...int) stackList[T] {
	if len(capacity) == 0 {
		return stackList[T]{topNode: nil, top: -1, capacity: StackMaxSize}
	}
	return stackList[T]{topNode: nil, top: -1, capacity: capacity[0]}
}

// IsEmpty returns true if Stack Top is -1
func (s *stackArray[T]) IsEmpty() bool {
	return s.top == -1
}

// IsEmpty returns true if Stack Top is -1
func (s *stackList[T]) IsEmpty() bool {
	return s.top == -1
}
