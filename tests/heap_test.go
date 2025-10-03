package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

func TestMinHeap(t *testing.T) {
	t.Run("Integer MinHeap", func(t *testing.T) {
		heap := linear.NewMinHeap[int]()

		// Test initial state
		if !heap.IsEmpty() {
			t.Error("New heap should be empty")
		}

		// Test Push and Size
		heap.Push(5)
		heap.Push(3)
		heap.Push(7)
		heap.Push(1)
		heap.Push(9)

		if heap.Size() != 5 {
			t.Errorf("Heap size should be 5, got %d", heap.Size())
		}

		// Test Peek (should return minimum)
		val, err := heap.Peek()
		if err != nil {
			t.Error("Peek should not return error for non-empty heap")
		}
		if val != 1 {
			t.Errorf("Peek should return 1 (minimum), got %d", val)
		}

		// Test Pop (should return elements in ascending order)
		expected := []int{1, 3, 5, 7, 9}
		for i, exp := range expected {
			val, err := heap.Pop()
			if err != nil {
				t.Errorf("Pop %d should not return error", i)
			}
			if val != exp {
				t.Errorf("Pop %d: expected %d, got %d", i, exp, val)
			}
		}

		// Test heap is empty after all pops
		if !heap.IsEmpty() {
			t.Error("Heap should be empty after popping all elements")
		}

		// Test Pop on empty heap
		_, err = heap.Pop()
		if err == nil {
			t.Error("Pop should return error for empty heap")
		}

		// Test Peek on empty heap
		_, err = heap.Peek()
		if err == nil {
			t.Error("Peek should return error for empty heap")
		}
	})

	t.Run("String MinHeap", func(t *testing.T) {
		heap := linear.NewMinHeap[string]()

		heap.Push("dog")
		heap.Push("cat")
		heap.Push("zebra")
		heap.Push("ant")

		val, err := heap.Peek()
		if err != nil || val != "ant" {
			t.Errorf("Peek should return 'ant', got '%s'", val)
		}

		// Test Clear
		heap.Clear()
		if !heap.IsEmpty() {
			t.Error("Heap should be empty after Clear")
		}
	})

	t.Run("MinHeap ToSlice", func(t *testing.T) {
		heap := linear.NewMinHeap[int]()
		heap.Push(5)
		heap.Push(3)
		heap.Push(7)

		slice := heap.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Slice length should be 3, got %d", len(slice))
		}

		// Verify it's a copy (modifying slice shouldn't affect heap)
		slice[0] = 999
		val, _ := heap.Peek()
		if val == 999 {
			t.Error("ToSlice should return a copy, not reference to internal array")
		}
	})
}

func TestMaxHeap(t *testing.T) {
	t.Run("Integer MaxHeap", func(t *testing.T) {
		heap := linear.NewMaxHeap[int]()

		// Test initial state
		if !heap.IsEmpty() {
			t.Error("New heap should be empty")
		}

		// Test Push and Size
		heap.Push(5)
		heap.Push(3)
		heap.Push(7)
		heap.Push(1)
		heap.Push(9)

		if heap.Size() != 5 {
			t.Errorf("Heap size should be 5, got %d", heap.Size())
		}

		// Test Peek (should return maximum)
		val, err := heap.Peek()
		if err != nil {
			t.Error("Peek should not return error for non-empty heap")
		}
		if val != 9 {
			t.Errorf("Peek should return 9 (maximum), got %d", val)
		}

		// Test Pop (should return elements in descending order)
		expected := []int{9, 7, 5, 3, 1}
		for i, exp := range expected {
			val, err := heap.Pop()
			if err != nil {
				t.Errorf("Pop %d should not return error", i)
			}
			if val != exp {
				t.Errorf("Pop %d: expected %d, got %d", i, exp, val)
			}
		}

		// Test heap is empty after all pops
		if !heap.IsEmpty() {
			t.Error("Heap should be empty after popping all elements")
		}

		// Test Pop on empty heap
		_, err = heap.Pop()
		if err == nil {
			t.Error("Pop should return error for empty heap")
		}

		// Test Peek on empty heap
		_, err = heap.Peek()
		if err == nil {
			t.Error("Peek should return error for empty heap")
		}
	})

	t.Run("Float MaxHeap", func(t *testing.T) {
		heap := linear.NewMaxHeap[float64]()

		heap.Push(3.14)
		heap.Push(2.71)
		heap.Push(9.81)
		heap.Push(1.41)

		val, err := heap.Peek()
		if err != nil || val != 9.81 {
			t.Errorf("Peek should return 9.81, got %f", val)
		}

		// Test Clear
		heap.Clear()
		if !heap.IsEmpty() {
			t.Error("Heap should be empty after Clear")
		}
	})

	t.Run("MaxHeap ToSlice", func(t *testing.T) {
		heap := linear.NewMaxHeap[int]()
		heap.Push(5)
		heap.Push(3)
		heap.Push(7)

		slice := heap.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Slice length should be 3, got %d", len(slice))
		}

		// Verify it's a copy (modifying slice shouldn't affect heap)
		slice[0] = -999
		val, _ := heap.Peek()
		if val == -999 {
			t.Error("ToSlice should return a copy, not reference to internal array")
		}
	})
}
