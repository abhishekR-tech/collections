package linear

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents a node in the deque
type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

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
func (d *Deque[T]) AddFirst(value T) {
	newNode := &Node[T]{
		value: value,
		prev:  nil,
		next:  d.head,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.head.prev = newNode
		d.head = newNode
	}
	d.length++
}

// AddLast adds an element to the end of the deque
func (d *Deque[T]) AddLast(value T) {
	newNode := &Node[T]{
		value: value,
		prev:  d.tail,
		next:  nil,
	}

	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
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

	value := d.head.value
	d.head = d.head.next
	d.length--

	if d.head == nil {
		d.tail = nil
	} else {
		d.head.prev = nil
	}

	return value, nil
}

// RemoveLast removes and returns the last element
func (d *Deque[T]) RemoveLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}

	value := d.tail.value
	d.tail = d.tail.prev
	d.length--

	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.next = nil
	}

	return value, nil
}

// PeekFirst returns the first element without removing it
func (d *Deque[T]) PeekFirst() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.head.value, nil
}

// PeekLast returns the last element without removing it
func (d *Deque[T]) PeekLast() (T, error) {
	var zero T
	if d.IsEmpty() {
		return zero, errors.New("deque is empty")
	}
	return d.tail.value, nil
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
	for i := 0; i < d.length; i++ {
		result[i] = current.value
		current = current.next
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
		sb.WriteString(fmt.Sprintf("%v", current.value))
		if current.next != nil {
			sb.WriteString(" ")
		}
		current = current.next
	}

	sb.WriteString("]")
	return sb.String()
}
