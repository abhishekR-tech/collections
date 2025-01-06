package linear

import "errors"

// A Stack that has generics so that it can be created
// for any given datatype
type Stack[T any] struct {
	items []T
}

// Constructor method
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Writing a function that checks a stack for empty
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

// Add an element into a stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Remove the last elemt in the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Empty Stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Look at the  "top" of the stack ,i.e., last element
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Empty Stack")
	}

	return s.items[len(s.items)-1], nil
}

// Convert the Slack into a slice for output
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}
