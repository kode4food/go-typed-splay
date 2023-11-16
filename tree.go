package splaytree

// SplayTree defines the splay tree type.
type SplayTree[Item any] struct {
	lessThan LessThan[Item]
	root     *node[Item]
}

// Compile time test if SplayTree fully implements Interface.
var _ Interface[any] = (*SplayTree[any])(nil)

// Clear all elements from the tree.
func (tree *SplayTree[_]) Clear() {
	tree.root = nil
}

// Count the number of elements in the tree.
func (tree *SplayTree[_]) Count() int {
	return tree.root.count()
}

// Height computes the height of the tree.
// This is the number of steps from
// the root to the farthest element.
// An empty tree has height -1.
func (tree *SplayTree[_]) Height() int {
	return -1 + tree.root.height()
}

// NonEmpty tests if the tree contains at least one element.
func (tree *SplayTree[_]) NonEmpty() bool {
	return tree.root != nil
}

// Root gives the element which is currently at the root of the tree. If the
// tree is empty then the zero value and false are returned.
func (tree *SplayTree[Item]) Root() (Item, bool) {
	if tree.root == nil {
		var zero Item
		return zero, false
	}
	return tree.root.item, true
}
