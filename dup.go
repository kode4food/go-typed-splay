package splaytree

// Duplicate clones a tree. The internal tree structure is
// duplicated, while the used items in the cloned tree
// are identical with the original tree.
func (tree *SplayTree[Item]) Duplicate() Interface[Item] {
	res := NewSplayTree[Item](tree.lt)
	res.root = tree.root.duplicate()
	return res
}
