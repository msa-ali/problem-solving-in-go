package tree

import "golang.org/x/exp/constraints"

type BST[T Number] struct {
	root *Node[T]
}

type Number interface {
	constraints.Integer | constraints.Float
}

func NewBST[T Number](root T) *BST[T] {
	return &BST[T]{
		root: NewNode(root),
	}
}

func (bst *BST[T]) Find(value T) bool {
	if bst == nil {
		return false
	}
	return walkAndSearch(bst.root, value)
}

func walkAndSearch[T Number](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if node.value > value {
		return walkAndSearch(node.left, value)
	}
	return walkAndSearch(node.right, value)
}
