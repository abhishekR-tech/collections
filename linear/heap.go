package linear

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// MinHeap represents a min-heap data structure
type MinHeap[T constraints.Ordered] struct {
	items []T
}

// MaxHeap represents a max-heap data structure
type MaxHeap[T constraints.Ordered] struct {
	items []T
}

// NewMinHeap creates a new MinHeap
func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		items: make([]T, 0),
	}
}

// NewMaxHeap creates a new MaxHeap
func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		items: make([]T, 0),
	}
}

// MinHeap methods

// IsEmpty checks if the heap is empty
func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of items in the heap
func (h *MinHeap[T]) Size() int {
	return len(h.items)
}

// Clear removes all items from the heap
func (h *MinHeap[T]) Clear() {
	h.items = make([]T, 0)
}

// Push adds an element to the heap
func (h *MinHeap[T]) Push(item T) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

// Pop removes and returns the minimum element from the heap
func (h *MinHeap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		return *new(T), errors.New("empty heap")
	}

	min := h.items[0]
	lastIdx := len(h.items) - 1
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return min, nil
}

// Peek returns the minimum element without removing it
func (h *MinHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		return *new(T), errors.New("empty heap")
	}
	return h.items[0], nil
}

// ToSlice returns a copy of the heap items
func (h *MinHeap[T]) ToSlice() []T {
	result := make([]T, len(h.items))
	copy(result, h.items)
	return result
}

// heapifyUp maintains heap property by moving element up
func (h *MinHeap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2
		if h.items[index] >= h.items[parentIdx] {
			break
		}
		h.items[index], h.items[parentIdx] = h.items[parentIdx], h.items[index]
		index = parentIdx
	}
}

// heapifyDown maintains heap property by moving element down
func (h *MinHeap[T]) heapifyDown(index int) {
	size := len(h.items)
	for {
		smallest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < size && h.items[leftChild] < h.items[smallest] {
			smallest = leftChild
		}
		if rightChild < size && h.items[rightChild] < h.items[smallest] {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}

// MaxHeap methods

// IsEmpty checks if the heap is empty
func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of items in the heap
func (h *MaxHeap[T]) Size() int {
	return len(h.items)
}

// Clear removes all items from the heap
func (h *MaxHeap[T]) Clear() {
	h.items = make([]T, 0)
}

// Push adds an element to the heap
func (h *MaxHeap[T]) Push(item T) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

// Pop removes and returns the maximum element from the heap
func (h *MaxHeap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		return *new(T), errors.New("empty heap")
	}

	max := h.items[0]
	lastIdx := len(h.items) - 1
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return max, nil
}

// Peek returns the maximum element without removing it
func (h *MaxHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		return *new(T), errors.New("empty heap")
	}
	return h.items[0], nil
}

// ToSlice returns a copy of the heap items
func (h *MaxHeap[T]) ToSlice() []T {
	result := make([]T, len(h.items))
	copy(result, h.items)
	return result
}

// heapifyUp maintains heap property by moving element up
func (h *MaxHeap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2
		if h.items[index] <= h.items[parentIdx] {
			break
		}
		h.items[index], h.items[parentIdx] = h.items[parentIdx], h.items[index]
		index = parentIdx
	}
}

// heapifyDown maintains heap property by moving element down
func (h *MaxHeap[T]) heapifyDown(index int) {
	size := len(h.items)
	for {
		largest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < size && h.items[leftChild] > h.items[largest] {
			largest = leftChild
		}
		if rightChild < size && h.items[rightChild] > h.items[largest] {
			largest = rightChild
		}

		if largest == index {
			break
		}

		h.items[index], h.items[largest] = h.items[largest], h.items[index]
		index = largest
	}
}
