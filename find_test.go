package splaytree

import "testing"

func TestMin(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.Check()
	if _, ok := tree.Min(); ok {
		t.Errorf("tree min !nil")
	}
	items := []int{3, 5, 7, 2, 4, 6, 8}
	min := 99
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.Check()
		if item < min {
			min = item
		}
		if m, ok := tree.Min(); ok && m != min {
			t.Errorf("tree min !%v", item)
		}
		tree.Check()
	}
}

func TestMax(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.Check()
	if _, ok := tree.Max(); ok {
		t.Errorf("tree max !nil")
	}
	items := []int{3, 5, 7, 2, 4, 6, 8}
	max := 0
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.Check()
		if max < item {
			max = item
		}
		if m, ok := tree.Max(); ok && m != max {
			t.Errorf("tree max !%v", item)
		}
		tree.Check()
	}
}

func TestLookup(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.Check()
	items := []int{3, 5, 7, 2, 4, 6, 8}
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.Check()
		if i, ok := tree.Lookup(item); !ok || i != item {
			t.Errorf("tree lookup !%v", item)
		}
		tree.Check()
	}
	miss := []int{1, 0, 9}
	for _, item := range miss {
		if i, ok := tree.Lookup(item); ok {
			t.Errorf("tree lookup !!%v %v", item, i)
		}
		tree.Check()
	}
}
