package splaytree

import "testing"

// Iterate over items in a tree while observing lower and upper bounds.
func TestRangeIterator(t *testing.T) {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	items := []int{2, 4, 6, 1, 5, 3, 0}
	tree.InsertAll(items)
	for lkup := range items {
		tree.Lookup(lkup)
		lower := 2
		upper := 4
		iter := tree.RangeIterator(lower, upper)
		for item, ok := iter(); ok; item, ok = iter() {
			if item < lower || upper < item {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = -10
		upper = -1
		iter = tree.RangeIterator(lower, upper)
		for item, ok := iter(); ok; item, ok = iter() {
			if item < lower || upper < item {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = -1
		upper = 3
		iter = tree.RangeIterator(lower, upper)
		for item, ok := iter(); ok; item, ok = iter() {
			if item < lower || upper < item {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = 3
		upper = 9
		iter = tree.RangeIterator(lower, upper)
		for item, ok := iter(); ok; item, ok = iter() {
			if item < lower || upper < item {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = 9
		upper = 29
		iter = tree.RangeIterator(lower, upper)
		for item, ok := iter(); ok; item, ok = iter() {
			if item < lower || upper < item {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
	}
}
