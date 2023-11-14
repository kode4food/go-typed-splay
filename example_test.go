package splaytree

import "fmt"

func Example() {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.Insert(2)
	tree.InsertAll([]int{1, 4, 0, 3})
	tree.Traverse(func(i int) { fmt.Printf("%v ", i) })
	iter := tree.Iterator()
	for i, ok := iter(); ok; i, ok = iter() {
		fmt.Printf("%v ", i)
	}
	// Output: 3 0 4 <nil> 1 2 1 2
}

// Iterate over items in a tree. Abort the iteration halfway
// and do a lookup to preserve optimal theoretical properties.
func ExampleSplayTree_Iterator() {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.InsertAll([]int{2, 4, 1, 3, 6, 0})
	iter := tree.Iterator()
	for item, ok := iter(); ok; item, ok = iter() {
		fmt.Printf("%v ", item)
		if item > 2 {
			tree.Lookup(item)
			break
		}
	}
	// Output: 0 1 2 3
}

// Iterate over items in a tree in reverse order.
// Aborting the iteration halfway requires doing a lookup on the last item.
func ExampleSplayTree_ReverseIterator() {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.InsertAll([]int{10, 99, 42, 0, 66, 5})
	iter := tree.ReverseIterator()
	for item, ok := iter(); ok; item, ok = iter() {
		fmt.Printf("%v ", item)
		if item == 42 {
			tree.Lookup(item)
			break
		}
	}
	// Output: 99 66 42
}

// Iterate over items in a tree while observing lower and upper bounds.
func ExampleSplayTree_RangeIterator() {
	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	tree.InsertAll([]int{2, 4, 6, 1, 5, 3, 0})
	iter := tree.RangeIterator(2, 4)
	var last int
	for item, ok := iter(); ok; item, ok = iter() {
		fmt.Printf("%v ", item)
		last = item
	}
	tree.Lookup(last)
	// Output: 2 3 4
}
