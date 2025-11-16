package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

// Stack Benchmarks

func BenchmarkStackPush(b *testing.B) {
	stack := linear.NewStack[int]()
	b.ResetTimer()
	for i := range b.N {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := linear.NewStack[int]()
	for i := range 10000 {
		stack.Push(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		stack.Pop()
		if stack.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				stack.Push(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkStackPeek(b *testing.B) {
	stack := linear.NewStack[int]()
	stack.Push(42)
	b.ResetTimer()
	for _ = range b.N {
		stack.Peek()
	}
}

// Queue Benchmarks

func BenchmarkQueueEnqueue(b *testing.B) {
	queue := linear.NewQueue[int]()
	b.ResetTimer()
	for i := range b.N {
		queue.Enqueue(i)
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	queue := linear.NewQueue[int]()
	for i := range 10000 {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		queue.Dequeue()
		if queue.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				queue.Enqueue(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkQueuePeek(b *testing.B) {
	queue := linear.NewQueue[int]()
	queue.Enqueue(42)
	b.ResetTimer()
	for _ = range b.N {
		queue.Peek()
	}
}

// Deque Benchmarks

func BenchmarkDequeAddFirst(b *testing.B) {
	deque := linear.NewDeque[int]()
	b.ResetTimer()
	for i := range b.N {
		deque.AddFirst(i)
	}
}

func BenchmarkDequeAddLast(b *testing.B) {
	deque := linear.NewDeque[int]()
	b.ResetTimer()
	for i := range b.N {
		deque.AddLast(i)
	}
}

func BenchmarkDequeRemoveFirst(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := range b.N {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		deque.RemoveFirst()
		if deque.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				deque.AddLast(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkDequeRemoveLast(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := range 10000 {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		deque.RemoveLast()
		if deque.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				deque.AddLast(j)
			}
			b.StartTimer()
		}
	}
}

// LinkedList Benchmarks

func BenchmarkLinkedListAppend(b *testing.B) {
	list := linear.NewLinkedList[int]()
	b.ResetTimer()
	for i := range b.N {
		list.Append(i)
	}
}

func BenchmarkLinkedListPrepend(b *testing.B) {
	list := linear.NewLinkedList[int]()
	b.ResetTimer()
	for i := range b.N {
		list.Prepend(i)
	}
}

// MinHeap Benchmarks

func BenchmarkMinHeapPush(b *testing.B) {
	heap := linear.NewMinHeap[int]()
	b.ResetTimer()
	for i := range b.N {
		heap.Push(i)
	}
}

func BenchmarkMinHeapPop(b *testing.B) {
	heap := linear.NewMinHeap[int]()
	for i := range 10000 {
		heap.Push(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		heap.Pop()
		if heap.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				heap.Push(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkMinHeapPeek(b *testing.B) {
	heap := linear.NewMinHeap[int]()
	heap.Push(42)
	b.ResetTimer()
	for _ = range b.N {
		heap.Peek()
	}
}

// MaxHeap Benchmarks

func BenchmarkMaxHeapPush(b *testing.B) {
	heap := linear.NewMaxHeap[int]()
	b.ResetTimer()
	for i := range b.N {
		heap.Push(i)
	}
}

func BenchmarkMaxHeapPop(b *testing.B) {
	heap := linear.NewMaxHeap[int]()
	for i := range 10000 {
		heap.Push(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		heap.Pop()
		if heap.IsEmpty() {
			b.StopTimer()
			for j := range 10000 {
				heap.Push(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkMaxHeapPeek(b *testing.B) {
	heap := linear.NewMaxHeap[int]()
	heap.Push(42)
	b.ResetTimer()
	for _ = range b.N {
		heap.Peek()
	}
}

// Conversion Benchmarks

func BenchmarkStackToSlice(b *testing.B) {
	stack := linear.NewStack[int]()
	for i := range b.N {
		stack.Push(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		stack.ToSlice()
	}
}

func BenchmarkQueueToSlice(b *testing.B) {
	queue := linear.NewQueue[int]()
	for i := range b.N {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		queue.ToSlice()
	}
}

func BenchmarkLinkedListToSlice(b *testing.B) {
	list := linear.NewLinkedList[int]()
	for i := range b.N {
		list.Append(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		list.ToSlice()
	}
}

func BenchmarkDequeToSlice(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := range b.N {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for _ = range b.N {
		deque.ToSlice()
	}
}
