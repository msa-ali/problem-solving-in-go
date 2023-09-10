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

func TestCompareBT(t *testing.T) {
	t1 := NewTree[int](7)
	t1.root.left = NewNode(23)
	t1.root.right = NewNode(3)
	t1.root.left.left = NewNode(5)
	t1.root.left.right = NewNode(4)
	t1.root.right.left = NewNode(18)
	t1.root.right.right = NewNode(21)
	t2 := NewTree[int](7)
	t2.root.left = NewNode(23)
	t2.root.right = NewNode(3)
	t2.root.left.left = NewNode(5)
	t2.root.left.right = NewNode(4)
	t2.root.right.left = NewNode(18)
	t2.root.right.right = NewNode(21)
	assert.True(t, Compare(t1, t2))
}

func TestCompareBT2(t *testing.T) {
	t1 := NewTree[int](7)
	t1.root.left = NewNode(23)
	t1.root.left.left = NewNode(5)
	t1.root.left.right = NewNode(4)
	t2 := NewTree[int](7)
	t2.root.left = NewNode(23)
	t2.root.right = NewNode(3)
	t2.root.left.left = NewNode(5)
	t2.root.left.right = NewNode(4)
	t2.root.right.left = NewNode(18)
	t2.root.right.right = NewNode(21)
	assert.False(t, Compare(t1, t2))
}

/*
			[8]
		   / 	\
		[3]	  	[10]
		/  \
	  [1]	[6]
	       /
		 [4]
*/
func TestBSTSearch(t *testing.T) {
	bst := NewBST[int](8)
	bst.root.left = NewNode(3)
	bst.root.right = NewNode(10)
	bst.root.left.left = NewNode(1)
	bst.root.left.right = NewNode(6)
	bst.root.left.right.left = NewNode(4)
	assert.True(t, bst.Find(4))
	assert.False(t, bst.Find(0))
	assert.False(t, bst.Find(11))
	assert.True(t, bst.Find(6))
	assert.True(t, bst.Find(10))
}
