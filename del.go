package splaytree

// DeleteRoot removes the element which is currently at the root.
// Return nil if the tree was empty, else the root item.
func (tree *SplayTree[Item]) DeleteRoot() (Item, bool) {
	node := tree.root
	if node == nil {
		var zero Item
		return zero, false
	}
	found := node.item
	if node.left == nil {
		tree.root = node.right
	} else if node.right == nil {
		tree.root = node.left
	} else {
		temp := node.right
		tree.root = node.left
		tree.splay(found)
		tree.root.right = temp
	}
	return found, true
}

// Delete an item from the tree.
// Return the deleted item if it was found, else nil.
func (tree *SplayTree[Item]) Delete(item Item) (Item, bool) {
	if tree.root == nil {
		var zero Item
		return zero, false
	}
	tree.splay(item)
	if tree.lt(item, tree.root.item) || tree.lt(tree.root.item, item) {
		var zero Item
		return zero, false
	}
	return tree.DeleteRoot()
}

// DeleteAll deletes all given items.
// Return the number of deleted items.
func (tree *SplayTree[Item]) DeleteAll(items []Item) int {
	num := 0
	for _, item := range items {
		if _, ok := tree.Delete(item); ok {
			num++
		}
	}
	return num
}

// DeleteMin deletes the smallest element from the tree.
// Return the smallest item if the tree was non-empty, else nil.
func (tree *SplayTree[Item]) DeleteMin() (Item, bool) {
	node := tree.root
	if node == nil {
		var zero Item
		return zero, false
	}
	if node.left == nil {
		tree.root = node.right
	} else {
		var parent = node
		node = node.left
		for node.left != nil {
			parent = node
			node = node.left
		}
		parent.left = node.right
		tree.splay(parent.item)
	}
	return node.item, true
}

// DeleteMax deletes the largest element from the tree.
// Return the largest item if the tree was non-empty, else nil.
func (tree *SplayTree[Item]) DeleteMax() (Item, bool) {
	node := tree.root
	if node == nil {
		var zero Item
		return zero, false
	}
	if node.right == nil {
		tree.root = node.left
	} else {
		var parent = node
		node = node.right
		for node.right != nil {
			parent = node
			node = node.right
		}
		parent.right = node.left
		tree.splay(parent.item)
	}
	return node.item, true
}
