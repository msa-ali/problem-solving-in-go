package tree

func LevelOrderBinaryTree(arr []int) *Tree[int] {
	tree := new(Tree[int])
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree[T any](arr []T, start int, size int) *Node[T] {
	curr := &Node[T]{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}
