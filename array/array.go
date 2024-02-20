package array // Package for array implementation

import (
	"errors"
	"fmt"
)

const ArrayMaxSize = 100 // Maximum size for the array

// array defines a fixed-size array data structure
type array[T float32 | float64 | int] struct {
	arr  [ArrayMaxSize]T // Array to hold elements
	size int             // Current number of elements in the array
}

// New creates a new instance of an array
func New[T float32 | float64 | int]() array[T] {
	return array[T]{size: 0} // Initialize with size 0
}

// Size returns the current size of the array
func (arr *array[T]) Size() int {
	return arr.size
}

// PushElement adds an element to the end of the array
func (arr *array[T]) PushElement(element T) error {
	if arr.size == ArrayMaxSize {
		return errors.New("Array is full")
	}

	arr.arr[arr.size] = element // Add element at the end
	arr.size++                  // Increment size
	return nil
}

// PopElement removes and returns the last element from the array
func (arr *array[T]) PopElement() (T, error) {
	if arr.size == 0 {
		return 0, errors.New("Array is empty")
	}

	arr.size-- // Decrement size before returning
	return arr.arr[arr.size], nil
}

// InsertElement inserts an element at a specific index in the array
func (arr *array[T]) InsertElement(element T, index int) error {
	if index < 0 || index > arr.size {
		return errors.New("Index out of bounds")
	}

	if arr.size == ArrayMaxSize {
		return errors.New("Array is full")
	}

	// Shift elements to the right to make space
	for i := arr.size - 1; i >= index; i-- {
		arr.arr[i+1] = arr.arr[i]
	}

	arr.arr[index] = element // Insert element at the index
	arr.size++               // Increment size
	return nil
}

// RemoveAtIndex removes the element at a specific index from the array
func (arr *array[T]) RemoveAtIndex(index int) error {
	if index < 0 || index >= arr.size {
		return errors.New("Index out of bounds")
	}

	if arr.size == 0 {
		return errors.New("Array is empty")
	}

	// Shift elements to the left to fill the gap
	for i := index; i < arr.size-1; i++ {
		arr.arr[i] = arr.arr[i+1]
	}

	arr.size-- // Decrement size
	return nil
}

// Get returns the element at a specific index from the array
func (arr *array[T]) Get(index int) (T, error) {
	if index < 0 || index >= arr.size {
		return 0, errors.New("Index out of bounds")
	}

	return arr.arr[index], nil
}

// IndexOf searches for an element in the array and returns its index
func (arr *array[T]) IndexOf(element T) (int, error) {
	for i := 0; i < arr.size; i++ {
		if arr.arr[i] == element {
			return i, nil
		}
	}

	return -1, errors.New("Element not found")
}

// PrintAll prints all elements of the array in a human-readable format
func (arr *array[T]) PrintAll() {
	fmt.Print("[ ")
	for i := 0; i < arr.size-1; i++ {
		fmt.Print(arr.arr[i], ", ")
	}
	fmt.Println(arr.arr[arr.size-1], "]")
}
