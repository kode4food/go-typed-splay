package splaytree

// Supremum finds the smallest element greater than a given.
// Return the supremum if present, else nil.
func (tree *SplayTree[Item]) Supremum(item Item) (Item, bool) {
	if tree.root == nil {
		var zero Item
		return zero, false
	}
	tree.splay(item)
	if tree.lessThan(item, tree.root.item) {
		return tree.root.item, true
	}
	if tree.root.right != nil {
		node := tree.root.right
		for node.left != nil {
			node = node.left
		}
		return node.item, true
	}
	var zero Item
	return zero, false
}

// Infimum finds the largest element smaller than a given.
// Return the infimum if present, else nil.
func (tree *SplayTree[Item]) Infimum(item Item) (Item, bool) {
	if tree.root == nil {
		var zero Item
		return zero, false
	}
	tree.splay(item)
	if tree.lessThan(tree.root.item, item) {
		return tree.root.item, true
	}
	if tree.root.left != nil {
		node := tree.root.left
		for node.right != nil {
			node = node.right
		}
		return node.item, true
	}
	var zero Item
	return zero, false
}
