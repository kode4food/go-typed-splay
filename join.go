package splaytree

// Join two trees with cost O(N + M), which is optimal.
// After the join all unique elements of both trees are
// in the first tree. The second tree becomes empty.
func (tree *SplayTree[Item]) Join(other Interface[Item]) {
	if other == nil {
		return
	}
	oak := other.(*SplayTree[Item])
	if oak == tree {
		return
	}
	tree.root = tree.join(tree.root, oak.root)
	oak.root = nil
}

func (tree *SplayTree[Item]) join(fir, oak *node[Item]) *node[Item] {
	if fir == nil {
		return oak
	} else if oak == nil {
		return fir
	}
	if tree.lt(fir.item, oak.item) {
		oak = tree.join(fir.right, oak)
		fir.right = nil
		oak.left = tree.join(fir, oak.left)
		return oak
	} else if tree.lt(oak.item, fir.item) {
		fir = tree.join(fir, oak.right)
		oak.right = nil
		fir.left = tree.join(fir.left, oak)
		return fir
	} else {
		fir.left = tree.join(fir.left, oak.left)
		fir.right = tree.join(fir.right, oak.right)
		return fir
	}
}
