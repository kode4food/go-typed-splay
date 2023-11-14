package splaytree

// Traverse a tree inorder with a function.
func (tree *SplayTree[Item]) Traverse(apply func(Item)) {
	tree.root.traverse(apply)
}

func (node *node[Item]) traverse(apply func(Item)) {
	if node != nil {
		node.left.traverse(apply)
		apply(node.item)
		node.right.traverse(apply)
	}
}
