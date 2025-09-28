package linear

import (
	"errors"
)

type LinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (ll *LinkedList[T]) Append(item T) {
	newNode := &Node[T]{
		Value: item,
		Prev:  ll.tail,
		Next:  nil,
	}

	if ll.tail != nil {
		ll.tail.Next = newNode
	} else {
		ll.head = newNode
	}

	ll.tail = newNode
	ll.length++
}

func (ll *LinkedList[T]) Prepend(item T) {
	newNode := &Node[T]{
		Value: item,
		Prev:  nil,
		Next:  ll.head,
	}

	if ll.head != nil {
		ll.head.Prev = newNode
	} else {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

func (ll *LinkedList[T]) Insert(index int, item T) {
	current := ll.head
	for _ = range index {
		current = current.Next
	}
	newNode := &Node[T]{
		Value: item,
		Prev:  current,
		Next:  current.Next,
	}
	if current.Next != nil {
		nextNode := current.Next
		nextNode.Prev = newNode
	}
	current.Next = newNode
	ll.length++
}

// Delete removes the element at the specified index
func (ll *LinkedList[T]) Delete(index int) (T, error) {
	var zero T

	// Bounds checking
	if index < 0 || index >= ll.length {
		return zero, errors.New("index out of bounds")
	}

	if ll.IsEmpty() {
		return zero, errors.New("list is empty")
	}

	var current *Node[T]
	var value T

	// Special case: delete head (index 0)
	if index == 0 {
		value = ll.head.Value
		ll.head = ll.head.Next
		if ll.head != nil {
			ll.head.Prev = nil
		} else {
			ll.tail = nil // List is now empty
		}
		ll.length--
		return value, nil
	}

	// Special case: delete tail (last element)
	if index == ll.length-1 {
		value = ll.tail.Value
		ll.tail = ll.tail.Prev
		if ll.tail != nil {
			ll.tail.Next = nil
		} else {
			ll.head = nil // List is now empty
		}
		ll.length--
		return value, nil
	}

	// General case: traverse to index and delete
	current = ll.head
	for range index {
		current = current.Next
	}

	value = current.Value

	// Update pointers to remove current node
	if current.Prev != nil {
		current.Prev.Next = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}

	ll.length--
	return value, nil
}

// Get retrieves the element at the specified index
func (ll *LinkedList[T]) Get(index int) (T, error) {
	var zero T

	// Bounds checking
	if index < 0 || index >= ll.length {
		return zero, errors.New("index out of bounds")
	}

	if ll.IsEmpty() {
		return zero, errors.New("list is empty")
	}
	var current *Node[T]

	// Optimization: traverse from head or tail based on which is closer
	if index < ll.length/2 {
		// Traverse from head
		current = ll.head
		for range index {
			current = current.Next
		}
	} else {
		// Traverse from tail
		current = ll.tail
		stepsFromTail := ll.length - 1 - index
		for range stepsFromTail {
			current = current.Prev
		}
	}
	return current.Value, nil
}

// Find searches for the first occurrence of a value and returns its index
// Returns -1 if not found
func (ll *LinkedList[T]) Find(value T, equal func(T, T) bool) int {
	if ll.IsEmpty() {
		return -1
	}

	current := ll.head

	// Modern approach: use range over int for indexing
	for i := range ll.length {
		if equal(current.Value, value) {
			return i
		}
		current = current.Next
	}

	return -1
}

func (ll *LinkedList[T]) Size() int {
	return ll.length
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.length == 0
}
