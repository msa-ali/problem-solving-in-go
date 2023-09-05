package tree

func (t *Tree[T]) PrintPreOrder() (res []T) {
	res = printPreOrder(t.root, res)
	return
}

func printPreOrder[T any](n *Node[T], result []T) []T {
	if n == nil {
		return result
	}
	result = append(result, n.value)
	if n.left != nil {
		result = printPreOrder(n.left, result)
	}
	if n.right != nil {
		result = printPreOrder(n.right, result)
	}
	return result
}
