package splaytree

type node[Item any] struct {
	left, right *node[Item]
	item        Item
}

func newNode[Item any](item Item) *node[Item] {
	return &node[Item]{item: item}
}

func (node *node[_]) count() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.count() + node.right.count()
}

func (node *node[_]) height() int {
	if node == nil {
		return 0
	}
	return 1 + max(node.left.height(), node.right.height())
}

func (node *node[Item]) duplicate() *node[Item] {
	if node == nil {
		return nil
	}
	res := newNode(node.item)
	res.left = node.left.duplicate()
	res.right = node.right.duplicate()
	return res
}
