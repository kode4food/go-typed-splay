package splaytree

import "testing"

func TestInsert(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	num := 0
	for _, item := range []int{6, 4, 2, 5, 3, 7, 0} {
		if tree.Insert(item) == false {
			t.Errorf("insert !%v", item)
		}
		num++
		if cnt := tree.Count(); cnt != num {
			t.Errorf("insert %v != %v", item, cnt)
		}
		if _, ok := tree.Lookup(item); !ok {
			t.Errorf("lookup after insert !%v", item)
		}
		tree.Check()
	}
}

func TestReplace(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	num := 0
	for _, item := range []int{6, 4, 2, 5, 3, 7, 0} {
		if tree.Replace(item) != false {
			t.Errorf("replace !%v false", item)
		}
		num++
		if cnt := tree.Count(); cnt != num {
			t.Errorf("insert %v != %v", item, cnt)
		}
		if _, ok := tree.Lookup(item); !ok {
			t.Errorf("lookup after replace !%v", item)
		}
		tree.Check()
	}
	for _, item := range []int{6, 4, 2, 5, 3, 7, 0} {
		if tree.Replace(item) != true {
			t.Errorf("replace !%v true", item)
		}
		if cnt := tree.Count(); cnt != num {
			t.Errorf("replace %v != %v", item, cnt)
		}
		if _, ok := tree.Lookup(item); !ok {
			t.Errorf("lookup after 2nd replace !%v", item)
		}
		tree.Check()
	}
}

func TestReplaceAll(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	items1 := []int{6, 4, 2, 5, 3, 7, 0}
	num := tree.ReplaceAll(items1)
	if num != 0 {
		t.Errorf("ReplaceAll %v != %v", num, 0)
	}
	items2 := []int{6, 4, 2, 5, 3, 7, 0}
	num = tree.ReplaceAll(items2)
	if num != len(items2) {
		t.Errorf("ReplaceAll %v != %v", num, len(items2))
	}
	tree.Check()
}
