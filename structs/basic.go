package structs

/**
This package is created for the basic structs
that is needed everywhere
*/

// Node represents a node
type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}
