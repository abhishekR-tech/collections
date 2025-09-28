package tests

import (
	"log"
	"strings"
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

// Helper function to create a populated list for testing
func createTestList() *linear.LinkedList[int] {
	list := linear.NewLinkedList[int]()
	for i := 1; i <= 5; i++ {
		list.Append(i * 10) // [10, 20, 30, 40, 50]
	}
	return list
}

// Helper function for integer equality
func intEqual(a, b int) bool {
	return a == b
}

// Helper function for string equality (case-insensitive)
func stringEqualIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}

// TestNewLinkedList tests the constructor
func TestNewLinkedList(t *testing.T) {
	list := linear.NewLinkedList[int]()

	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	if list.Size() != 0 {
		t.Errorf("New list size should be 0, got %d", list.Size())
	}
}

// TestLinkedListBasicOperations tests append, prepend, size, and empty checks
func TestLinkedListBasicOperations(t *testing.T) {
	list := linear.NewLinkedList[string]()

	// Test initial state
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	// Test append
	list.Append("first")
	if list.IsEmpty() {
		t.Error("List should not be empty after append")
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	// Test multiple appends
	list.Append("second")
	list.Append("third")
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
}

// TestLinkedListGet tests the Get method
func TestLinkedListGet(t *testing.T) {
	list := createTestList()

	tests := []struct {
		name        string
		index       int
		expected    int
		shouldError bool
	}{
		{"Get first element", 0, 10, false},
		{"Get middle element", 2, 30, false},
		{"Get last element", 4, 50, false},
		{"Get negative index", -1, 0, true},
		{"Get index too large", 5, 0, true},
		{"Get index way too large", 100, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Print(tt.index)
			log.Print(list.Size())
			value, err := list.Get(tt.index)

			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error for index %d, but got none", tt.index)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for index %d: %v", tt.index, err)
				}
				if value != tt.expected {
					t.Errorf("Expected %d, got %d", tt.expected, value)
				}
			}
		})
	}
}

// TestLinkedListGetEmpty tests Get on empty list
func TestLinkedListGetEmpty(t *testing.T) {
	list := linear.NewLinkedList[int]()

	_, err := list.Get(0)
	if err == nil {
		t.Error("Expected error when getting from empty list")
	}
}

// TestLinkedListGetBidirectionalOptimization tests that Get chooses optimal traversal direction
func TestLinkedListGetBidirectionalOptimization(t *testing.T) {
	// Create a larger list to test optimization
	list := linear.NewLinkedList[int]()
	for i := range 10 {
		list.Append(i)
	}

	// Test elements closer to head
	value, err := list.Get(1)
	if err != nil || value != 1 {
		t.Errorf("Expected 1, got %d with error %v", value, err)
	}

	// Test elements closer to tail
	value, err = list.Get(8)
	if err != nil || value != 8 {
		t.Errorf("Expected 8, got %d with error %v", value, err)
	}
}

// TestLinkedListDelete tests the Delete method
func TestLinkedListDelete(t *testing.T) {
	tests := []struct {
		name           string
		setupList      func() *linear.LinkedList[int]
		deleteIndex    int
		expectedValue  int
		expectedSize   int
		shouldError    bool
		remainingItems []int
	}{
		{
			name:           "Delete head from single element",
			setupList:      func() *linear.LinkedList[int] { l := linear.NewLinkedList[int](); l.Append(42); return l },
			deleteIndex:    0,
			expectedValue:  42,
			expectedSize:   0,
			shouldError:    false,
			remainingItems: []int{},
		},
		{
			name:           "Delete head from multiple elements",
			setupList:      createTestList,
			deleteIndex:    0,
			expectedValue:  10,
			expectedSize:   4,
			shouldError:    false,
			remainingItems: []int{20, 30, 40, 50},
		},
		{
			name:           "Delete tail from multiple elements",
			setupList:      createTestList,
			deleteIndex:    4,
			expectedValue:  50,
			expectedSize:   4,
			shouldError:    false,
			remainingItems: []int{10, 20, 30, 40},
		},
		{
			name:           "Delete middle element",
			setupList:      createTestList,
			deleteIndex:    2,
			expectedValue:  30,
			expectedSize:   4,
			shouldError:    false,
			remainingItems: []int{10, 20, 40, 50},
		},
		{
			name:           "Delete negative index",
			setupList:      createTestList,
			deleteIndex:    -1,
			expectedValue:  0,
			expectedSize:   5,
			shouldError:    true,
			remainingItems: []int{10, 20, 30, 40, 50},
		},
		{
			name:           "Delete index too large",
			setupList:      createTestList,
			deleteIndex:    10,
			expectedValue:  0,
			expectedSize:   5,
			shouldError:    true,
			remainingItems: []int{10, 20, 30, 40, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setupList()

			value, err := list.Delete(tt.deleteIndex)

			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error for index %d, but got none", tt.deleteIndex)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for index %d: %v", tt.deleteIndex, err)
				}
				if value != tt.expectedValue {
					t.Errorf("Expected deleted value %d, got %d", tt.expectedValue, value)
				}
			}

			// Check size
			if list.Size() != tt.expectedSize {
				t.Errorf("Expected size %d after deletion, got %d", tt.expectedSize, list.Size())
			}

			// Check remaining items
			for i, expected := range tt.remainingItems {
				actual, err := list.Get(i)
				if err != nil {
					t.Errorf("Error getting item at index %d: %v", i, err)
				}
				if actual != expected {
					t.Errorf("Expected item at index %d to be %d, got %d", i, expected, actual)
				}
			}
		})
	}
}

// TestLinkedListDeleteEmpty tests Delete on empty list
func TestLinkedListDeleteEmpty(t *testing.T) {
	list := linear.NewLinkedList[int]()

	_, err := list.Delete(0)
	if err == nil {
		t.Error("Expected error when deleting from empty list")
	}
}

// TestLinkedListFind tests the Find method with custom comparator
func TestLinkedListFind(t *testing.T) {
	list := createTestList()

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"Find first element", 10, 0},
		{"Find middle element", 30, 2},
		{"Find last element", 50, 4},
		{"Find non-existent element", 99, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := list.Find(tt.value, intEqual)
			if index != tt.expected {
				t.Errorf("Expected index %d, got %d", tt.expected, index)
			}
		})
	}
}

// TestLinkedListFindEmpty tests Find on empty list
func TestLinkedListFindEmpty(t *testing.T) {
	list := linear.NewLinkedList[int]()

	index := list.Find(42, intEqual)
	if index != -1 {
		t.Errorf("Expected -1 for empty list, got %d", index)
	}
}

// TestLinkedListFindWithStrings tests Find with string type and custom comparator
func TestLinkedListFindWithStrings(t *testing.T) {
	list := linear.NewLinkedList[string]()
	list.Append("Hello")
	list.Append("WORLD")
	list.Append("golang")

	// Test case-sensitive search
	index := list.Find("Hello", func(a, b string) bool { return a == b })
	if index != 0 {
		t.Errorf("Expected index 0 for case-sensitive 'Hello', got %d", index)
	}

	// Test case-insensitive search
	index = list.Find("world", stringEqualIgnoreCase)
	if index != 1 {
		t.Errorf("Expected index 1 for case-insensitive 'world', got %d", index)
	}

	// Test not found
	index = list.Find("python", func(a, b string) bool { return a == b })
	if index != -1 {
		t.Errorf("Expected -1 for 'python', got %d", index)
	}
}

// TestLinkedListFindMultipleOccurrences tests finding multiple occurrences
func TestLinkedListFindMultipleOccurrences(t *testing.T) {
	list := linear.NewLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(10)
	list.Append(30)
	list.Append(10)

	// Find first occurrence
	index := list.Find(10, intEqual)
	if index != 0 {
		t.Errorf("Expected first occurrence at index 0, got %d", index)
	}
}

// BenchmarkLinkedListGet benchmarks the Get operation
func BenchmarkLinkedListGet(b *testing.B) {
	list := linear.NewLinkedList[int]()
	for i := range 1000 {
		list.Append(i)
	}

	b.ResetTimer()
	for _ = range b.N {
		list.Get(500) // Get middle element
	}
}

// BenchmarkLinkedListDelete benchmarks the Delete operation
func BenchmarkLinkedListDelete(b *testing.B) {
	b.ResetTimer()
	for _ = range b.N {
		b.StopTimer()
		list := linear.NewLinkedList[int]()
		for j := range 100 {
			list.Append(j)
		}
		b.StartTimer()

		list.Delete(50) // Delete middle element
	}
}

// BenchmarkLinkedListFind benchmarks the Find operation
func BenchmarkLinkedListFind(b *testing.B) {
	list := linear.NewLinkedList[int]()
	for i := range 1000 {
		list.Append(i)
	}

	b.ResetTimer()
	for _ = range b.N {
		list.Find(500, intEqual) // Find middle element
	}
}
