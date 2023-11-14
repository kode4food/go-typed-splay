package splaytree

// Insert adds a new item to the tree, if it is not yet in the tree.
// Return true if the element was added, else false.
func (tree *SplayTree[Item]) Insert(item Item) bool {
	return tree.insertReplace(item, false)
}

// Replace an item if it was already in the tree.
// Otherwise insert the item as a new element.
// Return true if it was replaced, false if inserted.
func (tree *SplayTree[Item]) Replace(item Item) bool {
	return tree.insertReplace(item, true)
}

// InsertAll inserts a number of elements to the tree.
// Return the number of added items,
// i.e. the increase in size of the tree.
func (tree *SplayTree[Item]) InsertAll(items []Item) int {
	num := 0
	for _, item := range items {
		if tree.insertReplace(item, false) {
			num++
		}
	}
	return num
}

// ReplaceAll replaces or inserts a number of elements.
// Return the number of elements which replaced existing items.
func (tree *SplayTree[Item]) ReplaceAll(items []Item) int {
	if items == nil {
		panic("SplayTree ReplaceAll nil")
	}
	num := 0
	for _, item := range items {
		if tree.insertReplace(item, true) {
			num++
		}
	}
	return num
}

func (tree *SplayTree[Item]) insertReplace(item Item, replace bool) bool {
	if tree.root == nil {
		tree.root = newNode(item)
		return !replace
	}
	tree.splay(item)
	root := tree.root
	if tree.lt(item, root.item) {
		node := newNode(item)
		node.left = root.left
		node.right = root
		root.left = nil
		tree.root = node
		return !replace
	} else if tree.lt(root.item, item) {
		node := newNode(item)
		node.right = root.right
		node.left = root
		root.right = nil
		tree.root = node
		return !replace
	} else if replace {
		root.item = item
	}
	return replace
}
