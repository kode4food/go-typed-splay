package splaytree

// Min gives the smallest element in the tree. If the tree is empty then the
// zero value and false are returned.
func (tree *SplayTree[Item]) Min() (Item, bool) {
	node := tree.root
	if node == nil {
		var zero Item
		return zero, false
	}
	for node.left != nil {
		node = node.left
	}
	item := node.item
	tree.splay(item)
	return item, true
}

// Max gives the largest element in the tree. If the tree is empty then the
// zero value and false are returned
func (tree *SplayTree[Item]) Max() (Item, bool) {
	node := tree.root
	if node == nil {
		var zero Item
		return zero, false
	}
	for node.right != nil {
		node = node.right
	}
	item := node.item
	tree.splay(item)
	return item, true
}

// Lookup an item and return the found item. If the tree is empty then the zero
// value and false are returned.
func (tree *SplayTree[Item]) Lookup(item Item) (Item, bool) {
	if tree.root == nil {
		var zero Item
		return zero, false
	}
	tree.splay(item)
	if tree.lessThan(item, tree.root.item) || tree.lessThan(tree.root.item, item) {
		var zero Item
		return zero, false
	}
	return tree.root.item, true
}
