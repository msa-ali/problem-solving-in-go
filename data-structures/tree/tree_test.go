package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompleteBinaryTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	expected := []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7}
	res := tree.PrintPreOrder()
	assert.Equal(t, expected, res)
}
