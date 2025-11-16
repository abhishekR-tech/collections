# Collections

A comprehensive Go package providing generic data structures using Go 1.18+ generics.

## Features

- **Type-safe**: All data structures use Go generics for compile-time type safety
- **Zero dependencies**: Pure Go implementation (except for constraints package)
- **Well-tested**: Comprehensive test coverage
- **Documented**: Full godoc documentation for all exported functions

## Available Data Structures

### Linear Data Structures

- **Stack**: LIFO (Last In First Out) data structure
- **Queue**: FIFO (First In First Out) data structure
- **Deque**: Double-ended queue supporting operations at both ends
- **LinkedList**: Doubly linked list with bidirectional traversal
- **MinHeap**: Min-heap for efficient minimum element retrieval
- **MaxHeap**: Max-heap for efficient maximum element retrieval

## Installation

```bash
go get github.com/abhishekR-tech/collections
```

## Usage Examples

### Stack

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    stack := linear.NewStack[int]()

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    top, _ := stack.Peek()
    fmt.Println(top) // Output: 3

    value, _ := stack.Pop()
    fmt.Println(value) // Output: 3

    fmt.Println(stack.Size()) // Output: 2
}
```

### Queue

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    queue := linear.NewQueue[string]()

    queue.Enqueue("first")
    queue.Enqueue("second")
    queue.Enqueue("third")

    first, _ := queue.Peek()
    fmt.Println(first) // Output: first

    value, _ := queue.Dequeue()
    fmt.Println(value) // Output: first

    fmt.Println(queue.Size()) // Output: 2
}
```

### Deque

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    deque := linear.NewDeque[int]()

    deque.AddFirst(2)
    deque.AddFirst(1)
    deque.AddLast(3)
    deque.AddLast(4)

    fmt.Println(deque.String()) // Output: [1 2 3 4]

    first, _ := deque.RemoveFirst()
    fmt.Println(first) // Output: 1

    last, _ := deque.RemoveLast()
    fmt.Println(last) // Output: 4
}
```

### LinkedList

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    list := linear.NewLinkedList[int]()

    list.Append(1)
    list.Append(2)
    list.Append(3)
    list.Prepend(0)

    value, _ := list.Get(2)
    fmt.Println(value) // Output: 2

    // Find element
    equalFunc := func(a, b int) bool { return a == b }
    index := list.Find(2, equalFunc)
    fmt.Println(index) // Output: 2

    fmt.Println(list.Size()) // Output: 4
}
```

### MinHeap

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    heap := linear.NewMinHeap[int]()

    heap.Push(5)
    heap.Push(2)
    heap.Push(8)
    heap.Push(1)

    min, _ := heap.Peek()
    fmt.Println(min) // Output: 1

    value, _ := heap.Pop()
    fmt.Println(value) // Output: 1

    value, _ = heap.Pop()
    fmt.Println(value) // Output: 2
}
```

### MaxHeap

```go
package main

import (
    "fmt"
    "github.com/abhishekR-tech/collections/linear"
)

func main() {
    heap := linear.NewMaxHeap[int]()

    heap.Push(5)
    heap.Push(2)
    heap.Push(8)
    heap.Push(1)

    max, _ := heap.Peek()
    fmt.Println(max) // Output: 8

    value, _ := heap.Pop()
    fmt.Println(value) // Output: 8

    value, _ = heap.Pop()
    fmt.Println(value) // Output: 5
}
```

## Running Tests

```bash
go test ./tests/... -v
```

## Requirements

- Go 1.18 or higher (for generics support)

## License

This project is open source and available for use.
