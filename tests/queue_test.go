package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

func TestQueue_Integer(t *testing.T) {
	q := linear.NewQueue[int]()

	// Test initial state
	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	if q.Size() != 0 {
		t.Error("new queue should have size 0")
	}

	// Test Enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.IsEmpty() {
		t.Error("queue should not be empty after enqueuing")
	}
	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}

	// Test Peek
	item, err := q.Peek()
	if err != nil {
		t.Error("unexpected error on peek:", err)
	}
	if item != 1 {
		t.Errorf("expected peek to return 1, got %d", item)
	}

	// Test Dequeue
	item, err = q.Dequeue()
	if err != nil {
		t.Error("unexpected error on dequeue:", err)
	}
	if item != 1 {
		t.Errorf("expected dequeue to return 1, got %d", item)
	}
	if q.Size() != 2 {
		t.Errorf("expected size 2 after dequeue, got %d", q.Size())
	}

	// Test Clear
	q.Clear()
	if !q.IsEmpty() {
		t.Error("queue should be empty after clear")
	}

	// Test Dequeue on empty queue
	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}

func TestQueue_String(t *testing.T) {
	q := linear.NewQueue[string]()

	// Test string operations
	q.Enqueue("hello")
	q.Enqueue("world")

	item, err := q.Dequeue()
	if err != nil {
		t.Error("unexpected error on dequeue:", err)
	}
	if item != "hello" {
		t.Errorf("expected 'hello', got '%s'", item)
	}
}

func TestQueue_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := linear.NewQueue[Person]()

	// Test struct operations
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}

	q.Enqueue(p1)
	q.Enqueue(p2)

	item, err := q.Peek()
	if err != nil {
		t.Error("unexpected error on peek:", err)
	}
	if item.Name != "Alice" {
		t.Errorf("expected name 'Alice', got '%s'", item.Name)
	}
}
