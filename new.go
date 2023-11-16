package splaytree

type LessThan[Item any] func(l, r Item) bool

// NewSplayTree constructs a new empty splay tree.
func NewSplayTree[Item any](lt LessThan[Item]) *SplayTree[Item] {
	return &SplayTree[Item]{
		lessThan: lt,
	}
}
