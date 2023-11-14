package splaytree

import "testing"

func TestNewSplayTree(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	if tree.root != nil {
		t.Errorf("new root != nil")
	}
}
