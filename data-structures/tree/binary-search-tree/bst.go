package binarysearchtree

import (
	"fmt"
)

type Node struct {
	val         int
	left, right *Node
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type BinarySearchTree struct {
	root  *Node
	count int
}

func NewBST(val int) *BinarySearchTree {
	return &BinarySearchTree{
		root:  NewNode(val),
		count: 1,
	}
}

func insertRecursively(node *Node, val int) {
	if val > node.val {
		if node.right == nil {
			node.right = NewNode(val)
			return
		}
		insertRecursively(node.right, val)
		return
	}
	if node.left == nil {
		node.left = NewNode(val)
		return
	}
	insertRecursively(node.left, val)
}

func insertIteratively(node *Node, val int) {
	var parent *Node
	current := node
	for current != nil {
		parent = current
		if val > current.val {
			current = current.right
		} else {
			current = current.left
		}
	}
	if val > parent.val {
		parent.right = NewNode(val)
	} else {
		parent.left = NewNode(val)
	}

}

func (bst *BinarySearchTree) Insert(val int, iterative bool) {
	if iterative {
		insertIteratively(bst.root, val)
	} else {
		insertRecursively(bst.root, val)
	}
	bst.count++
}

func preorder(root *Node, result *[]int) []int {
	if root == nil {
		return *result
	}
	*result = append(*result, root.val)
	preorder(root.left, result)
	preorder(root.right, result)
	return *result
}

func inorder(root *Node, result *[]int) []int {
	if root == nil {
		return *result
	}
	inorder(root.left, result)
	*result = append(*result, root.val)
	inorder(root.right, result)
	return *result
}

func postOrder(root *Node, result *[]int) []int {
	if root == nil {
		return *result
	}
	postOrder(root.left, result)
	postOrder(root.right, result)
	*result = append(*result, root.val)
	return *result
}

func (bst *BinarySearchTree) PreOrder() []int {
	result := make([]int, 0, bst.count)
	preorder(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) InOrder() []int {
	result := make([]int, 0, bst.count)
	inorder(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) PostOrder() []int {
	result := make([]int, 0, bst.count)
	postOrder(bst.root, &result)
	return result
}

func (bst *BinarySearchTree) Has(val int) bool {
	current := bst.root
	for current != nil {
		if val == current.val {
			return true
		}
		if val < current.val {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

func findMaxInSubtree(root *Node) (maxNode, previous *Node) {
	previous = nil
	current := root
	for current.right != nil {
		previous = current
		current = current.right
	}
	maxNode = current
	return
}

func findMinInSubtree(root *Node) (minNode, previous *Node) {
	previous = nil
	current := root
	for current.left != nil {
		previous = current
		current = current.left
	}
	minNode = current
	return
}

func (bst *BinarySearchTree) max() (maxNode, previous *Node) {
	return findMaxInSubtree(bst.root)
}

func (bst *BinarySearchTree) min() (minNode, previous *Node) {
	return findMinInSubtree(bst.root)
}

func DeleteNodeFromBST(root *Node, val int) error {
	if root == nil {
		return fmt.Errorf("can't delete from empty BST")
	}
	curr := root
	var parent *Node
	for curr != nil {
		if curr.val == val {
			// perform deletion
			// case 1: leaf node - no child
			if curr.left == nil && curr.right == nil {
				if parent.left == curr {
					parent.left = nil
				} else {
					parent.right = nil
				}
			} else if curr.left != nil && curr.right != nil {
				// case 2: both children
				var nodeToReplace, previousNode *Node
				if val <= root.val {
					nodeToReplace, previousNode = findMaxInSubtree(root.left)
				} else {
					nodeToReplace, previousNode = findMinInSubtree(root.right)
				}
				if previousNode.left == nodeToReplace {
					previousNode.left = nil
				} else {
					previousNode.right = nil
				}
				curr.val = nodeToReplace.val

			} else {
				// case 3: one child
				temp := curr.left
				if curr.right != nil {
					temp = curr.right
				}
				if parent.left == curr {
					parent.left = temp
				} else {
					parent.right = temp
				}
			}
			return nil
		} else {
			parent = curr
			if val < curr.val {
				curr = curr.left
			} else {
				curr = curr.right
			}
		}
	}
	return fmt.Errorf("can't find node with value %d in BST", val)
}

func (bst *BinarySearchTree) Delete(val int) error {
	return DeleteNodeFromBST(bst.root, val)
}
