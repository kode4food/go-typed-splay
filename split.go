package splaytree

// Split a tree in two at a given boundary.
// Return two trees as a pair (left, right),
// where left contains all items less than the boundary
// and right contains only items at or beyond the boundary.
func (tree *SplayTree[Item]) Split(item Item) (*SplayTree[Item], *SplayTree[Item]) {
	more := NewSplayTree[Item](tree.lt)
	if tree.root == nil {
		return tree, more
	}
	tree.splay(item)
	if tree.lt(tree.root.item, item) {
		more.root = tree.root.right
		tree.root.right = nil
	} else {
		more.root = tree.root
		tree.root = more.root.left
		more.root.left = nil
	}
	return tree, more
}
