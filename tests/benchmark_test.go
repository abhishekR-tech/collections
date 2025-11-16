package tests

import (
	"testing"

	"github.com/abhishekR-tech/collections/linear"
)

// Stack Benchmarks

func BenchmarkStackPush(b *testing.B) {
	stack := linear.NewStack[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := linear.NewStack[int]()
	for i := 0; i < 10000; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
		if stack.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
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
	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

// Queue Benchmarks

func BenchmarkQueueEnqueue(b *testing.B) {
	queue := linear.NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	queue := linear.NewQueue[int]()
	for i := 0; i < 10000; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
		if queue.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
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
	for i := 0; i < b.N; i++ {
		queue.Peek()
	}
}

// Deque Benchmarks

func BenchmarkDequeAddFirst(b *testing.B) {
	deque := linear.NewDeque[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.AddFirst(i)
	}
}

func BenchmarkDequeAddLast(b *testing.B) {
	deque := linear.NewDeque[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.AddLast(i)
	}
}

func BenchmarkDequeRemoveFirst(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := 0; i < 10000; i++ {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.RemoveFirst()
		if deque.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
				deque.AddLast(j)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkDequeRemoveLast(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := 0; i < 10000; i++ {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.RemoveLast()
		if deque.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
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
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

func BenchmarkLinkedListPrepend(b *testing.B) {
	list := linear.NewLinkedList[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Prepend(i)
	}
}

// MinHeap Benchmarks

func BenchmarkMinHeapPush(b *testing.B) {
	heap := linear.NewMinHeap[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(i)
	}
}

func BenchmarkMinHeapPop(b *testing.B) {
	heap := linear.NewMinHeap[int]()
	for i := 0; i < 10000; i++ {
		heap.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Pop()
		if heap.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
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
	for i := 0; i < b.N; i++ {
		heap.Peek()
	}
}

// MaxHeap Benchmarks

func BenchmarkMaxHeapPush(b *testing.B) {
	heap := linear.NewMaxHeap[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(i)
	}
}

func BenchmarkMaxHeapPop(b *testing.B) {
	heap := linear.NewMaxHeap[int]()
	for i := 0; i < 10000; i++ {
		heap.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Pop()
		if heap.IsEmpty() {
			b.StopTimer()
			for j := 0; j < 10000; j++ {
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
	for i := 0; i < b.N; i++ {
		heap.Peek()
	}
}

// Conversion Benchmarks

func BenchmarkStackToSlice(b *testing.B) {
	stack := linear.NewStack[int]()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.ToSlice()
	}
}

func BenchmarkQueueToSlice(b *testing.B) {
	queue := linear.NewQueue[int]()
	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.ToSlice()
	}
}

func BenchmarkLinkedListToSlice(b *testing.B) {
	list := linear.NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.ToSlice()
	}
}

func BenchmarkDequeToSlice(b *testing.B) {
	deque := linear.NewDeque[int]()
	for i := 0; i < 1000; i++ {
		deque.AddLast(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.ToSlice()
	}
}
