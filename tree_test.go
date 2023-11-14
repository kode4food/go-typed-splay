package splaytree

import "testing"

func TestClear(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.InsertAll([]int{6, 4, 3, 1})
	tree.Check()
	tree.Clear()
	if tree.root != nil {
		t.Errorf("root != nil")
	}
}

func TestCount(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree.Count() != 0 {
		t.Errorf("count !0")
	}
	tree.InsertAll([]int{6, 4, 3, 1})
	tree.Check()
	if tree.Count() != 4 {
		t.Errorf("count !4")
	}
	tree.DeleteAll([]int{3, 2})
	if tree.Count() != 3 {
		t.Errorf("count !3")
	}
	tree.Clear()
	if tree.Count() != 0 {
		t.Errorf("count !0")
	}
}

func TestHeight(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree.Height() != -1 {
		t.Errorf("height !-1")
	}
	tree.Insert(3)
	if tree.Height() != 0 {
		t.Errorf("height !0")
	}
	tree.Insert(3)
	if tree.Height() != 0 {
		t.Errorf("height !0")
	}
	tree.Insert(1)
	if h := tree.Height(); h != 1 {
		t.Errorf("height !1 %v %v", h, tree.Count())
	}
	tree.Insert(2)
	if h := tree.Height(); h != 1 {
		t.Errorf("height !1 %v %v", h, tree.Count())
	}
	tree.Insert(9)
	if h := tree.Height(); h != 2 && h != 3 {
		t.Errorf("height !<2 %v %v", h, tree.Count())
	}
	tree.Check()
}

func TestNonEmpty(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree.NonEmpty() != false {
		t.Errorf("NonEmpty !false")
	}
	tree.Insert(2)
	if tree.NonEmpty() != true {
		t.Errorf("NonEmpty !true")
	}
	tree.Clear()
	if tree.NonEmpty() {
		t.Errorf("NonEmpty clear")
	}
}

func TestRoot(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if _, ok := tree.Root(); ok {
		t.Errorf("Root !false")
	}
	tree.Insert(2)
	if _, ok := tree.Root(); !ok {
		t.Errorf("Root !true")
	}
	tree.Clear()
	if _, ok := tree.Root(); ok {
		t.Errorf("Root !false")
	}
}
