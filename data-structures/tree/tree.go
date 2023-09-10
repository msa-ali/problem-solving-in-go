package tree

type Node[T any] struct {
	value       T
	left, right *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

type Tree[T any] struct {
	root *Node[T]
}

func NewTree[T any](root T) *Tree[T] {
	return &Tree[T]{
		root: NewNode(root),
	}
}
