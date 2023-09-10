package tree

func (t *Tree[T]) PreOrder() (res []T) {
	preOrderWalk(t.root, &res)
	return
}

func (t *Tree[T]) InOrder() (res []T) {
	inOrderWalk(t.root, &res)
	return
}

func (t *Tree[T]) PostOrder() (res []T) {
	postOrderWalk(t.root, &res)
	return
}

func (t *Tree[T]) LevelOrder() (res []T) {
	levelOrderWalk(t.root, &res)
	return
}

func preOrderWalk[T any](root *Node[T], res *[]T) {
	if root == nil {
		return
	}
	*res = append(*res, root.value)
	preOrderWalk(root.left, res)
	preOrderWalk(root.right, res)
}

func inOrderWalk[T any](root *Node[T], res *[]T) {
	if root == nil {
		return
	}
	inOrderWalk(root.left, res)
	*res = append(*res, root.value)
	inOrderWalk(root.right, res)
}

func postOrderWalk[T any](root *Node[T], res *[]T) {
	if root == nil {
		return
	}
	postOrderWalk(root.left, res)
	postOrderWalk(root.right, res)
	*res = append(*res, root.value)
}

func levelOrderWalk[T any](root *Node[T], res *[]T) {
	var queue []*Node[T]
	queue = append(queue, root)
	for len(queue) > 0 {
		var node *Node[T]
		node, queue = queue[0], queue[1:]
		*res = append(*res, node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
