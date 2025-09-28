package linear

import (
	"errors"
	"fmt"
	"strings"
)

//  Node represents a node in the deque
//
//	type Node[T any] struct {
//		value T
//		prev  *Node[T]
//		next  *Node[T]
//	}

// Deque represents a double-ended queue
type Deque[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// NewDeque creates a new empty deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// AddFirst adds an element to the front of the deque
func (d *Deque[T]) AddFirst(item T) {
	newNode := &Node[T]{
		Value: item,
		Prev:  nil,
		Next:  d.head,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.head.Prev = newNode
		d.head = newNode
	}
	d.length++
}

// AddLast adds an element to the end of the deque
func (d *Deque[T]) AddLast(item T) {
	newNode := &Node[T]{
		Value: item,
		Prev:  d.tail,
		Next:  nil,
	}

	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.Next = newNode
		d.tail = newNode
	}
	d.length++
}

// RemoveFirst removes and returns the first element
func (d *Deque[T]) RemoveFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}

	item := d.head.Value
	d.head = d.head.Next
	d.length--

	if d.head == nil {
		d.tail = nil
	} else {
		d.head.Prev = nil
	}

	return item, nil
}

// RemoveLast removes and returns the last element
func (d *Deque[T]) RemoveLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}

	value := d.tail.Value
	d.tail = d.tail.Prev
	d.length--

	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.Next = nil
	}

	return value, nil
}

// PeekFirst returns the first element without removing it
func (d *Deque[T]) PeekFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.head.Value, nil
}

// PeekLast returns the last element without removing it
func (d *Deque[T]) PeekLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.tail.Value, nil
}

// IsEmpty returns true if the deque has no elements
func (d *Deque[T]) IsEmpty() bool {
	return d.length == 0
}

// Size returns the number of elements in the deque
func (d *Deque[T]) Size() int {
	return d.length
}

// Clear removes all elements from the deque
func (d *Deque[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.length = 0
}

// ToSlice converts the deque to a slice
func (d *Deque[T]) ToSlice() []T {
	result := make([]T, d.length)
	current := d.head
	for i := range d.length {
		result[i] = current.Value
		current = current.Next
	}
	return result
}

// FromSlice creates a new deque from a slice
func FromSlice[T any](slice []T) *Deque[T] {
	deque := NewDeque[T]()
	for _, value := range slice {
		deque.AddLast(value)
	}
	return deque
}

// String returns a string representation of the deque
func (d *Deque[T]) String() string {
	if d.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	current := d.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.Value))
		if current.Next != nil {
			sb.WriteString(" ")
		}
		current = current.Next
	}

	sb.WriteString("]")
	return sb.String()
}
