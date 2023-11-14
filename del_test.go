package splaytree

import "testing"

func TestDeleteEmpty(t *testing.T) {
	tree := NewSplayTree[any](func(l, r any) bool {
		return false
	})
	if _, ok := tree.DeleteRoot(); ok {
		t.Errorf("DeleteRoot nil")
	}
	if _, ok := tree.DeleteMax(); ok {
		t.Errorf("DeleteMax nil")
	}
	if _, ok := tree.DeleteMin(); ok {
		t.Errorf("DeleteMin nil")
	}
	if tree.DeleteAll([]any{}) != 0 {
		t.Errorf("DeleteAll nil")
	}
}

func TestDelete(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	items := []int{6, 4, 2, 5, 3, 7, 0}
	tree.InsertAll(items)
	tree.Check()
	if n := tree.Count(); n != len(items) {
		t.Errorf("delete all count != %v", n)
	}
	if i, ok := tree.DeleteMax(); !ok || i != 7 || tree.Count() != len(items)-1 {
		t.Errorf("delete !max")
	}
	tree.Check()
	if i, ok := tree.DeleteMin(); !ok || i != 0 || tree.Count() != len(items)-2 {
		t.Errorf("delete !min")
	}
	tree.Check()
	r := tree.root.item
	if i, ok := tree.DeleteRoot(); !ok || i != r || tree.Count() != len(items)-3 {
		t.Errorf("delete !root")
	}
	tree.Check()
	if n := tree.DeleteAll(items); n != len(items)-3 || tree.Count() != 0 {
		t.Errorf("delete !all")
	}
	tree.Check()
}

func TestDeleteMin(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if _, ok := tree.DeleteMin(); ok {
		t.Errorf("DeleteMin nil")
	}
	tree.Insert(3)
	if i, ok := tree.DeleteMin(); !ok || i != 3 || tree.Count() != 0 {
		t.Errorf("DeleteMin 3")
	}
	tree.Insert(3)
	tree.Insert(2)
	if i, ok := tree.DeleteMin(); !ok || i != 2 || tree.Count() != 1 {
		t.Errorf("DeleteMin 3 2")
	}
	tree.Insert(2)
	tree.Insert(3)
	if i, ok := tree.DeleteMin(); !ok || i != 2 || tree.Count() != 1 {
		t.Errorf("DeleteMin 2 3")
	}
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	if i, ok := tree.DeleteMin(); !ok || i != 1 || tree.Count() != 2 {
		t.Errorf("DeleteMin 3 2 1")
	}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	if i, ok := tree.DeleteMin(); !ok || i != 1 || tree.Count() != 2 {
		t.Errorf("DeleteMin 2 3 1")
	}
}

func TestDeleteMax(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if _, ok := tree.DeleteMax(); ok {
		t.Errorf("DeleteMax nil")
	}
	tree.Insert(3)
	if i, ok := tree.DeleteMax(); !ok || i != 3 || tree.Count() != 0 {
		t.Errorf("DeleteMax 3")
	}
	tree.Insert(3)
	tree.Insert(2)
	if i, ok := tree.DeleteMax(); !ok || i != 3 || tree.Count() != 1 {
		t.Errorf("DeleteMax 3 2")
	}
	tree.Insert(2)
	tree.Insert(3)
	if i, ok := tree.DeleteMax(); !ok || i != 3 || tree.Count() != 1 {
		t.Errorf("DeleteMax 2 3")
	}
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	if i, ok := tree.DeleteMax(); !ok || i != 3 || tree.Count() != 2 {
		t.Errorf("DeleteMax 3 2 1")
	}
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	if i, ok := tree.DeleteMax(); !ok || i != 3 || tree.Count() != 2 {
		t.Errorf("DeleteMax 2 3 1")
	}
}
