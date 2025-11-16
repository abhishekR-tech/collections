package linear

import (
	"errors"
	"fmt"
	"strings"
)

// Queue represents a FIFO (First In First Out) data structure
type Queue[T any] struct {
	items []T
}

// NewQueue creates and returns a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item from the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("Empty Queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the first item in the queue without removing it
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("Empty Queue")
	}
	return q.items[0], nil
}

// IsEmpty returns true if the queue has no items
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue
func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}

// ToSlice returns a copy of the queue as a slice
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.items))
	copy(result, q.items)
	return result
}

// FromQueueSlice creates a new queue from a slice
func FromQueueSlice[T any](slice []T) *Queue[T] {
	queue := NewQueue[T]()
	queue.items = make([]T, len(slice))
	copy(queue.items, slice)
	return queue
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	for i, item := range q.items {
		sb.WriteString(fmt.Sprintf("%v", item))
		if i < len(q.items)-1 {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
