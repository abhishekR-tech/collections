package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

func TestStack(t *testing.T) {
	// Test integer stack
	t.Run("Integer Stack", func(t *testing.T) {
		stack := linear.NewStack[int]()

		// Test initial state
		if !stack.IsEmpty() {
			t.Error("New stack should be empty")
		}

		// Test Push and Size
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.Size() != 3 {
			t.Errorf("Stack size should be 3, got %d", stack.Size())
		}

		// Test Peek
		val, err := stack.Peek()
		if err != nil {
			t.Error("Peek should not return error for non-empty stack")
		}
		if val != 3 {
			t.Errorf("Peek should return 3, got %d", val)
		}

		// Test Pop
		val, err = stack.Pop()
		if err != nil {
			t.Error("Pop should not return error for non-empty stack")
		}
		if val != 3 {
			t.Errorf("Pop should return 3, got %d", val)
		}

		// Test ToSlice
		slice := stack.ToSlice()
		if len(slice) != 2 {
			t.Errorf("Slice length should be 2, got %d", len(slice))
		}
		if slice[0] != 1 || slice[1] != 2 {
			t.Error("Slice contents are incorrect")
		}

		// Test Clear
		stack.Clear()
		if !stack.IsEmpty() {
			t.Error("Stack should be empty after Clear")
		}

		// Test Pop on empty stack
		_, err = stack.Pop()
		if err == nil {
			t.Error("Pop should return error for empty stack")
		}
	})

	// Test string stack
	t.Run("String Stack", func(t *testing.T) {
		stack := linear.NewStack[string]()
		stack.Push("hello")
		stack.Push("world")

		val, err := stack.Pop()
		if err != nil || val != "world" {
			t.Error("String stack not working correctly")
		}
	})
}
