package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompleteBinaryTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	expected := []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7}
	res := tree.PreOrder()
	assert.Equal(t, expected, res)
}

/*
			  [7]
		   /	  \
		[23]	  [3]
		/  \	   /  \
	   [5]	[4]	  [18] [21]
*/
func TestTreeTraversal(t *testing.T) {
	tree := NewTree[int](7)
	tree.root.left = NewNode(23)
	tree.root.right = NewNode(3)
	tree.root.left.left = NewNode(5)
	tree.root.left.right = NewNode(4)
	tree.root.right.left = NewNode(18)
	tree.root.right.right = NewNode(21)
	res := tree.PreOrder()
	assert.Equal(t, []int{7, 23, 5, 4, 3, 18, 21}, res)
	res = tree.InOrder()
	assert.Equal(t, []int{5, 23, 4, 7, 18, 3, 21}, res)
	res = tree.PostOrder()
	assert.Equal(t, []int{5, 4, 23, 18, 21, 3, 7}, res)
	res = tree.LevelOrder()
	assert.Equal(t, []int{7, 23, 3, 5, 4, 18, 21}, res)
}
