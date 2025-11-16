package linear

// Container represents the basic interface for all collection types
type Container[T any] interface {
	// Size returns the number of elements in the container
	Size() int

	// IsEmpty returns true if the container has no elements
	IsEmpty() bool

	// Clear removes all elements from the container
	Clear()
}

// Stringable represents types that can be converted to string representation
type Stringable interface {
	// String returns a string representation of the data structure
	String() string
}

// Sliceable represents types that can be converted to/from slices
type Sliceable[T any] interface {
	// ToSlice returns a copy of the data structure as a slice
	ToSlice() []T
}

// Collection represents a generic collection with common operations
type Collection[T any] interface {
	Container[T]
	Stringable
	Sliceable[T]
}
