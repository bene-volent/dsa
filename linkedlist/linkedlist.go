package linkedlist

import (
	"errors"
	"fmt"
)

// Node structures for different linked list types
// ----------------------------------------------

// Node for a singly linked list
type unidirectionalNode[T any] struct {
	Next *unidirectionalNode[T] // Pointer to the next node in the list
	Val  T                      // Value stored in the node
}

// Node for a doubly linked list
type bidirectionalNode[T any] struct {
	Prev *bidirectionalNode[T] // Pointer to the previous node
	Next *bidirectionalNode[T] // Pointer to the next node
	Val  T                     // Value stored in the node
}

type node[T any] interface {
	val() T
}

// Interface for general linked list operations
// ------------------------------------------

type LinkedList[T int | float32 | float64] interface {
	// Insertion operations
	InsertAtBeginning(val T) error
	InsertAtEnd(val T) error
	InsertAtPosition(val T, pos int) error

	// Deletion operations
	DeleteFromBeginning() (T, error)
	DeleteFromEnd() (T, error)
	DeleteAtPosition(pos int) (T, error)

	// Traversal operation
	Traverse(operation func(T))

	// Search operation
	Search(val T) (bool, node[T])
}

// Singly Linked List Implementation
// --------------------------------

// SinglyLinkedList struct with head pointer and length
type SinglyLinkedList[T int | float32 | float64] struct {
	head   *unidirectionalNode[T] // Pointer to the first node in the list
	length int                    // Number of nodes in the list
}

// NewSLL returns a new Singly Linked List
func NewSLL[T int | float32 | float64]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{}
}

// Traversal function to visit each node and apply a given operation
// Time complexity: O(n)
func (l *SinglyLinkedList[T]) Traverse(operation func(T)) {
	current := l.head
	for current != nil {
		operation(current.Val) // Apply the operation to the current node's value
		current = current.Next
	}
}

// Insertion Operations
// -------------------

// InsertAtBeginning inserts a new node at the beginning of the list
// Time complexity: O(1)
func (l *SinglyLinkedList[T]) InsertAtBeginning(val T) error {
	newNode := &unidirectionalNode[T]{l.head, val}
	l.head = newNode
	l.length++
	return nil
}

// InsertAtEnd inserts a new node at the end of the list
// Time complexity: O(n)
func (l *SinglyLinkedList[T]) InsertAtEnd(val T) error {
	if l.head == nil {
		return errors.New("Cannot insert at end of empty list")
	}

	newNode := &unidirectionalNode[T]{nil, val}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.length++
	return nil
}

// InsertAtPosition inserts a new node at a specific position (0-based indexing)
//
// Time complexity: O(n)
func (l *SinglyLinkedList[T]) InsertAtPosition(val T, pos int) error {
	if pos < 0 || pos > l.length {
		return errors.New("Invalid position for insertion")
	}

	// Handle insertion at the beginning for efficiency
	if pos == 0 {
		return l.InsertAtBeginning(val)
	}

	if pos == l.length-1 {
		return l.InsertAtEnd(val)
	}
	// Create the new node to insert
	newNode := &unidirectionalNode[T]{nil, val}

	// Traverse to the node before the insertion position
	current := l.head
	for i := 1; i < pos; i++ { // Start from 1 since we already checked for pos = 0
		current = current.Next
	}

	// Insert the new node between the current node and its next node
	newNode.Next = current.Next
	current.Next = newNode

	// Increment the list length
	l.length++

	return nil
}

// DeleteFromBeginning deletes the node from the beginning of the singly linked list
//
// Time Complexity: O(1)
func (l *SinglyLinkedList[T]) DeleteFromBeginning() (T, error) {

	// Checks if the list is empty
	if l.length == 0 {
		return 0, errors.New("Cannot delete from an empty list!")
	}

	// Stores the value of the head to return
	val := l.head.Val

	// Deletes the current from the list
	l.head = l.head.Next
	l.length--

	return val, nil
}

// DeleteFromBeginning deletes the node from the end of the singly linked list
//
// Time Complexity: O(n)
func (l *SinglyLinkedList[T]) DeleteFromEnd() (T, error) {

	// Checks if the list is empty
	if l.length == 0 {
		return 0, errors.New("Cannot delete from an empty list!")
	}

	if l.length == 1 {
		return l.DeleteFromBeginning()
	}

	// Traverse to the before the tail of the linked list
	curr := l.head
	currNext := curr.Next
	for curr.Next != nil && currNext.Next != nil { // Starts from the head till the node before tail as tails next
		curr = curr.Next     // Till Node befor the tail
		currNext = curr.Next // Till the tail.
	}

	// Remove the tail from the list
	curr.Next = nil
	l.length--

	return currNext.Val, nil
}

// DeleteAtPosition deletes the node at the specified position from the singly linked list.
//
// Time Complexity: O(n)
func (l *SinglyLinkedList[T]) DeleteAtPosition(pos int) (T, error) {
	// Check for invalid positions and handle special cases efficiently.
	if pos < 0 || pos >= l.length {
		return 0, errors.New("invalid position for deletion")
	} else if pos == 0 {
		return l.DeleteFromBeginning()
	} else if pos == l.length-1 {
		return l.DeleteFromEnd()
	}

	// Traverse to the node before the one to be deleted.
	curr := l.head
	for i := 0; i < pos-1; i++ {
		curr = curr.Next
	}

	// Store the value of the node to be deleted.
	val := curr.Next.Val
	// Bypass the deleted node by linking the previous node to the next one.
	curr.Next = curr.Next.Next
	// Update the list length.
	l.length--

	// Return the deleted value and nil error.
	return val, nil

}

// Search searches for a given element in the singly linked list.
//
// Time Complexity : O(n)
func (l *SinglyLinkedList[T]) Search(element T) (bool, *unidirectionalNode[T]) {
	// Start searching from the head of the list.
	curr := l.head

	// Iterate through the list until the element is found or the end is reached.
	for curr != nil {
		if curr.Val == element {
			// Element found! Return true and the node.
			return true, curr
		}
		curr = curr.Next
	}

	// Element not found. Return false and nil.
	return false, nil
}

type DoublyLinkedList[T int | float32 | float64] struct {
	head   *bidirectionalNode[T]
	tail   *bidirectionalNode[T]
	length int
}

// NewSLL returns a new Singly Linked List
func NewDLL[T int | float32 | float64]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

// Traversal function to visit each node and apply a given operation
// Time complexity: O(n)
func (l *DoublyLinkedList[T]) Traverse(operation func(T)) {
	current := l.head
	for current != nil {
		operation(current.Val) // Apply the operation to the current node's value
		fmt.Printf("%v : %p\n", current, current)
		current = current.Next
	}
}

// Insertion Operations
// -------------------

// InsertAtBeginning inserts a new node at the beginning of the list
// Time complexity: O(1)
func (l *DoublyLinkedList[T]) InsertAtBeginning(val T) error {
	newNode := &bidirectionalNode[T]{Next: l.head, Prev: nil, Val: val}
	l.head = newNode
	if l.length == 0 {
		l.tail = newNode
	}

	l.length++
	return nil
}

// I'll do this later
// -------------------------------------------------------

// // InsertAtEnd inserts a new node at the end of the list
// // Time complexity: O(n)
// func (l *DoublyLinkedList[T]) InsertAtEnd(val T) error {
// 	newNode := &bidirectionalNode[T]{Next: nil, Prev: l.tail, Val: val}
// 	l.tail = newNode

// 	if l.length == 0 {
// 		l.head = newNode
// 		return nil
// 	}
// 	l.tail.Prev.Next =

// 	// fmt.Println(l.head, l.tail)
// 	l.length++
// 	return nil
// }

// // InsertAtPosition inserts a new node at a specific position (0-based indexing)
// //
// // Time complexity: O(n)
// func (l *DoublyLinkedList[T]) InsertAtPosition(val T, pos int) error {
// 	if pos < 0 || pos > l.length {
// 		return errors.New("Invalid position for insertion")
// 	}

// 	// Handle insertion at the beginning for efficiency
// 	if pos == 0 {
// 		return l.InsertAtBeginning(val)
// 	}
// 	if pos == l.length-1 {
// 		return l.InsertAtEnd(val)
// 	}

// 	// Create the new node to insert
// 	newNode := &bidirectionalNode[T]{nil, nil, val}

// 	// Traverse to the node before the insertion position
// 	current := l.head
// 	for i := 1; i < pos; i++ { // Start from 1 since we already checked for pos = 0
// 		current = current.Next
// 	}

// 	// Insert the new node between the current node and its next node
// 	newNode.Next = current.Next
// 	newNode.Prev = current
// 	current.Next = newNode
// 	newNode.Next.Prev = newNode

// 	// Increment the list length
// 	l.length++

// 	return nil
// }

// // // DeleteFromBeginning deletes the node from the beginning of the singly linked list
// // //
// // // Time Complexity: O(1)
// // func (l *SinglyLinkedList[T]) DeleteFromBeginning() (T, error) {

// // 	// Checks if the list is empty
// // 	if l.length == 0 {
// // 		return 0, errors.New("Cannot delete from an empty list!")
// // 	}

// // 	// Stores the value of the head to return
// // 	val := l.head.Val

// // 	// Deletes the current from the list
// // 	l.head = l.head.Next
// // 	l.length--

// // 	return val, nil
// // }

// // // DeleteFromBeginning deletes the node from the end of the singly linked list
// // //
// // // Time Complexity: O(n)
// // func (l *SinglyLinkedList[T]) DeleteFromEnd() (T, error) {

// // 	// Checks if the list is empty
// // 	if l.length == 0 {
// // 		return 0, errors.New("Cannot delete from an empty list!")
// // 	}

// // 	if l.length == 1 {
// // 		return l.DeleteFromBeginning()
// // 	}

// // 	// Traverse to the before the tail of the linked list
// // 	curr := l.head
// // 	currNext := curr.Next
// // 	for curr.Next != nil && currNext.Next != nil { // Starts from the head till the node before tail as tails next
// // 		curr = curr.Next     // Till Node befor the tail
// // 		currNext = curr.Next // Till the tail.
// // 	}

// // 	// Remove the tail from the list
// // 	curr.Next = nil
// // 	l.length--

// // 	return currNext.Val, nil
// // }

// // // DeleteAtPosition deletes the node at the specified position from the singly linked list.
// // //
// // // Time Complexity: O(n)
// // func (l *SinglyLinkedList[T]) DeleteAtPosition(pos int) (T, error) {
// // 	// Check for invalid positions and handle special cases efficiently.
// // 	if pos < 0 || pos >= l.length {
// // 		return 0, errors.New("invalid position for deletion")
// // 	} else if pos == 0 {
// // 		return l.DeleteFromBeginning()
// // 	} else if pos == l.length-1 {
// // 		return l.DeleteFromEnd()
// // 	}

// // 	// Traverse to the node before the one to be deleted.
// // 	curr := l.head
// // 	for i := 0; i < pos-1; i++ {
// // 		curr = curr.Next
// // 	}

// // 	// Store the value of the node to be deleted.
// // 	val := curr.Next.Val
// // 	// Bypass the deleted node by linking the previous node to the next one.
// // 	curr.Next = curr.Next.Next
// // 	// Update the list length.
// // 	l.length--

// // 	// Return the deleted value and nil error.
// // 	return val, nil

// // }

// // // Search searches for a given element in the singly linked list.
// // //
// // // Time Complexity : O(n)
// // func (l *SinglyLinkedList[T]) Search(element T) (bool, *unidirectionalNode[T]) {
// // 	// Start searching from the head of the list.
// // 	curr := l.head

// // 	// Iterate through the list until the element is found or the end is reached.
// // 	for curr != nil {
// // 		if curr.Val == element {
// // 			// Element found! Return true and the node.
// // 			return true, curr
// // 		}
// // 		curr = curr.Next
// // 	}

// // 	// Element not found. Return false and nil.
// // 	return false, nil
// // }
