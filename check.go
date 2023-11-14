package splaytree

// Check verifies that the tree
// is a proper binary search tree.
func (tree *SplayTree[_]) Check() {
	if !tree.isBinarySearchTree() {
		panic("IsBinarySearchTree failed")
	}
}

// Return true if tree is a proper binary search tree.
func (tree *SplayTree[_]) isBinarySearchTree() bool {
	if tree.root == nil {
		return true
	}
	_, _, ok := tree.isBinarySearchNode(tree.root)
	return ok
}

// Return the minimum and maximum element at this node.
func (tree *SplayTree[Item]) isBinarySearchNode(node *node[Item]) (Item, Item, bool) {
	min, max := node.item, node.item
	if node.left != nil {
		lmin, lmax, ok := tree.isBinarySearchNode(node.left)
		if !ok || !tree.lt(lmax, min) {
			return lmin, lmax, false
		}
		min = lmin
	}
	if node.right != nil {
		rmin, rmax, ok := tree.isBinarySearchNode(node.right)
		if !ok || !tree.lt(max, rmin) {
			return rmin, rmax, false
		}
		max = rmax
	}
	return min, max, true
}
