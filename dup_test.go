package splaytree

import "testing"

func TestDuplicate(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	items := []int{3, 5, 7, 2, 4, 6, 8}
	tree.InsertAll(items)
	tree.Check()
	dup := tree.Duplicate().(*SplayTree[int])
	dup.Check()
	if tree.Count() != dup.Count() {
		t.Errorf("dup count")
	}
	if tree.Height() != dup.Height() {
		t.Errorf("dup height")
	}
	for _, item := range items {
		if _, ok := tree.Lookup(item); !ok {
			t.Errorf("tree lookup %v", item)
		}
		if _, ok := dup.Lookup(item); !ok {
			t.Errorf("dup lookup %v", item)
		}
		tree.Check()
		dup.Check()
	}
}
