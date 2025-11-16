package linear

import (
	"errors"
	"fmt"
	"strings"
)

// Stack represents a LIFO (Last In First Out) data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates and returns a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// IsEmpty returns true if the stack has no items
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack
func (s *Stack[T]) Clear() {
	s.items = make([]T, 0)
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Empty Stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Empty Stack")
	}

	return s.items[len(s.items)-1], nil
}

// ToSlice returns a copy of the stack as a slice
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// FromStackSlice creates a new stack from a slice
func FromStackSlice[T any](slice []T) *Stack[T] {
	stack := NewStack[T]()
	stack.items = make([]T, len(slice))
	copy(stack.items, slice)
	return stack
}

// String returns a string representation of the stack (top to bottom)
func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	// Display from top to bottom (reverse order)
	for i := len(s.items) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%v", s.items[i]))
		if i > 0 {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
