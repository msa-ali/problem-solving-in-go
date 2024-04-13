package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
		    6
		/		\
		4		9
	   / \	   / \
	  2  5    8   12
	  			  /  \
			     10  14
*/
func CreateTestBST() *BinarySearchTree {
	bst := NewBST(6)
	bst.Insert(4, true)
	bst.Insert(9, false)
	bst.Insert(2, true)
	bst.Insert(5, true)
	bst.Insert(8, false)
	bst.Insert(12, false)
	bst.Insert(10, false)
	bst.Insert(14, true)
	return bst
}

func TestBST(t *testing.T) {
	bst := CreateTestBST()
	assert.Equal(t, bst.root.val, 6)
	assert.Equal(t, 4, bst.root.left.val)
	assert.Equal(t, 9, bst.root.right.val)
	assert.Equal(t, 2, bst.root.left.left.val)
	assert.Equal(t, 5, bst.root.left.right.val)
	assert.Equal(t, 8, bst.root.right.left.val)
	assert.Equal(t, 12, bst.root.right.right.val)
	assert.Equal(t, 10, bst.root.right.right.left.val)
	assert.Equal(t, 14, bst.root.right.right.right.val)
	assert.Equal(t, 9, bst.count)
	assert.Equal(t, []int{6, 4, 2, 5, 9, 8, 12, 10, 14}, bst.PreOrder())
	assert.Equal(t, []int{2, 4, 5, 6, 8, 9, 10, 12, 14}, bst.InOrder())
	assert.Equal(t, []int{2, 5, 4, 8, 10, 14, 12, 9, 6}, bst.PostOrder())
	assert.Equal(t, true, bst.Has(12))
	assert.Equal(t, true, bst.Has(6))
	assert.Equal(t, false, bst.Has(21))
	maxNode, _ := bst.max()
	assert.Equal(t, 14, maxNode.val)
	minNode, _ := bst.min()
	assert.Equal(t, 2, minNode.val)
}

func TestDeleteNodeFromBST(t *testing.T) {
	bst := CreateTestBST()
	assert.Equal(t, []int{2, 4, 5, 6, 8, 9, 10, 12, 14}, bst.InOrder())
	bst.Delete(14)
	assert.Equal(t, []int{2, 4, 5, 6, 8, 9, 10, 12}, bst.InOrder())
	bst.Delete(12)
	assert.Equal(t, []int{2, 4, 5, 6, 8, 9, 10}, bst.InOrder())
	bst = CreateTestBST()
	bst.Delete(4)
	assert.Equal(t, []int{2, 5, 6, 8, 9, 10, 12, 14}, bst.InOrder())
	bst.Delete(9)
	assert.Equal(t, []int{2, 5, 6, 8, 10, 12, 14}, bst.InOrder())
}

/*
		    6
		/		\
		4		9
	   / \	   / \
	  2  5    8   12
	  			  /  \
			     10  14

	        6
		/		\
		5		8
	   / \	   / \
	  2  nil  nil 12
	  			  /  \
			     10  14

		    5
		/		\
		4		6
	   / \	   / \
	  2  nil  8   12
	  			  /  \
			     10  14
*/
