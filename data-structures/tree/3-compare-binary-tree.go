package tree

func Compare[T comparable](a *Tree[T], b *Tree[T]) bool {
	if a == nil && b == nil {
		return true
	}
	return walkAndCompare(a.root, b.root)
}

// DFS preserves the shape of the traversal
func walkAndCompare[T comparable](node1 *Node[T], node2 *Node[T]) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil || node1.value != node2.value {
		return false
	}
	return walkAndCompare(node1.left, node2.left) && walkAndCompare(node1.right, node2.right)
}
