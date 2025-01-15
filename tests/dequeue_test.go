package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

func TestDeque(t *testing.T) {
	t.Run("new deque should be empty", func(t *testing.T) {
		deque := linear.NewDeque[int]()
		if !deque.IsEmpty() {
			t.Error("new deque should be empty")
		}
		if deque.Size() != 0 {
			t.Error("new deque should have size 0")
		}
	})

	t.Run("add and remove first", func(t *testing.T) {
		deque := linear.NewDeque[int]()
		deque.AddFirst(1)
		deque.AddFirst(2)

		if deque.Size() != 2 {
			t.Error("deque should have size 2")
		}

		val, err := deque.RemoveFirst()
		if err != nil || val != 2 {
			t.Error("expected 2 from RemoveFirst")
		}

		val, err = deque.RemoveFirst()
		if err != nil || val != 1 {
			t.Error("expected 1 from RemoveFirst")
		}

		if !deque.IsEmpty() {
			t.Error("deque should be empty after removing all elements")
		}
	})

	t.Run("add and remove last", func(t *testing.T) {
		deque := linear.NewDeque[int]()
		deque.AddLast(1)
		deque.AddLast(2)

		val, err := deque.RemoveLast()
		if err != nil || val != 2 {
			t.Error("expected 2 from RemoveLast")
		}

		val, err = deque.RemoveLast()
		if err != nil || val != 1 {
			t.Error("expected 1 from RemoveLast")
		}
	})

	t.Run("peek operations", func(t *testing.T) {
		deque := linear.NewDeque[int]()
		deque.AddFirst(1)
		deque.AddLast(2)

		val, err := deque.PeekFirst()
		if err != nil || val != 1 {
			t.Error("expected 1 from PeekFirst")
		}

		val, err = deque.PeekLast()
		if err != nil || val != 2 {
			t.Error("expected 2 from PeekLast")
		}
	})

	t.Run("to and from slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		deque := linear.FromSlice(slice)

		if deque.Size() != len(slice) {
			t.Error("deque size should match slice length")
		}

		newSlice := deque.ToSlice()
		for i, v := range slice {
			if newSlice[i] != v {
				t.Errorf("expected %v at index %d, got %v", v, i, newSlice[i])
			}
		}
	})

	t.Run("clear operation", func(t *testing.T) {
		deque := linear.NewDeque[int]()
		deque.AddFirst(1)
		deque.AddFirst(2)
		deque.Clear()

		if !deque.IsEmpty() {
			t.Error("deque should be empty after clear")
		}

		if deque.Size() != 0 {
			t.Error("deque size should be 0 after clear")
		}
	})
}
