package tree

type Node[T any] struct {
	value       T
	left, right *Node[T]
}

type Tree[T any] struct {
	root *Node[T]
}
