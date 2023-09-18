package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

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

func (bst *BST[T]) Insert(value T) {
	n := NewNode(value)
	if bst.root == nil {
		bst.root = n
		return
	}
	walkAndInsert(bst.root, value)
}

func walkAndInsert[T Number](node *Node[T], value T) {
	if node.value >= value {
		if node.left == nil {
			node.left = NewNode(value)
			return
		}
		walkAndInsert(node.left, value)
		return
	} else {
		if node.right == nil {
			node.right = NewNode(value)
			return
		}
		walkAndInsert(node.right, value)
		return
	}
}

func (bst *BST[T]) Delete(value T) (*Node[T], error) {
	if bst.root == nil {
		return nil, fmt.Errorf("binary search tree is not initialised")
	}
	return walkAndDelete(bst.root, value)
}
func walkAndDelete[T Number](node *Node[T], value T) (*Node[T], error) {
	curr := node
	var prev *Node[T]
	for curr != nil {
		if curr.value > value {
			prev = curr
			curr = curr.left
			continue
		}
		if curr.value < value {
			prev = curr
			curr = curr.right
			continue
		}
		// when curr.value == value
		deleteNode(curr, prev)
		return curr, nil
	}
	return nil, fmt.Errorf("unable to find node with value %v", node.value)
}

func deleteNode[T Number](curr *Node[T], prev *Node[T]) {
	// case 1: both child are nil
	if curr.left == nil && curr.right == nil {
		if prev.left == curr {
			prev.left = nil
		} else {
			prev.right = nil
		}
		return
	}
	// case 2: one of the child is nil
	if curr.left == nil {
		if prev.left == curr {
			prev.left = curr.right
		} else {
			prev.right = curr.right
		}
		return
	}
	if curr.right == nil {
		if prev.left == curr {
			prev.left = curr.left
		} else {
			prev.right = curr.left
		}
		return
	}
	// case 3: when both child are not nil
	maxNode, prevNode := getMaxNode(curr.left)
	deleteNode(maxNode, prevNode)
	curr.value = maxNode.value
}

func getMaxNode[T Number](node *Node[T]) (*Node[T], *Node[T]) {
	var max T
	var prev *Node[T]
	var maxNode *Node[T]
	var prevNode *Node[T]
	for node != nil {
		if node.value > max {
			max = node.value
			prevNode = prev
			maxNode = node
		}
		prev = node
		node = node.right
	}
	return maxNode, prevNode
}
