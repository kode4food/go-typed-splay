package splaytree

import "testing"

func TestInsert(t *testing.T) {
	tree := NewSplayTree()
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	num := 0
	for _, item := range []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)} {
		if tree.Insert(item) == false {
			t.Errorf("insert !%v", item)
		}
		num++
		if cnt := tree.Count(); cnt != num {
			t.Errorf("insert %v != %v", item, cnt)
		}
		if _, bo := tree.Lookup(item); bo == false {
			t.Errorf("lookup after insert !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
}

func TestReplace(t *testing.T) {
	tree := NewSplayTree()
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	num := 0
	for _, item := range []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)} {
		if tree.Replace(item) != false {
			t.Errorf("replace !%v false", item)
		}
		num++
		if cnt := tree.Count(); cnt != num {
			t.Errorf("insert %v != %v", item, cnt)
		}
		if _, bo := tree.Lookup(item); bo == false {
			t.Errorf("lookup after replace !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
	for _, item := range []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)} {
		if tree.Replace(item) != true {
			t.Errorf("replace !%v true", item)
		}
		if cnt := tree.Count(); cnt != num {
			t.Errorf("replace %v != %v", item, cnt)
		}
		if _, bo := tree.Lookup(item); bo == false {
			t.Errorf("lookup after 2nd replace !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
}

func TestReplaceAll(t *testing.T) {
	tree := NewSplayTree()
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	items1 := []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)}
	num := tree.ReplaceAll(items1)
	if num != 0 {
		t.Errorf("ReplaceAll %v != %v", num, 0)
	}
	items2 := []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)}
	num = tree.ReplaceAll(items2)
	if num != len(items2) {
		t.Errorf("ReplaceAll %v != %v", num, len(items2))
	}
	tree.CheckBinarySearchTree()
}
