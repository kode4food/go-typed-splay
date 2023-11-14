package splaytree

import "testing"
import "time"
import "math/rand"

func TestSplit(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []int {
		nums := rand.Perm(upper)
		items := make([]int, upper)
		for i := 0; i < upper; i++ {
			items[i] = nums[i]
		}
		return items
	}
	lessThan := func(bound int) func(int) {
		return func(item int) {
			if bound < item {
				t.Errorf("split %v is more than %v", item, bound)
			}
		}
	}
	atLeast := func(bound int) func(int) {
		return func(item int) {
			if item < bound {
				t.Errorf("split %v is less than %v", item, bound)
			}
		}
	}

	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	lef, rig := tree.Split(4)
	if lef.NonEmpty() || rig.NonEmpty() {
		t.Errorf("split nonempty")
	}

	for k := 1; k < 20; k++ {
		items := perm(k)
		for i := 0; i <= k; i++ {
			tree := NewSplayTree(func(l, r int) bool {
				return l < r
			})
			tree.InsertAll(items)
			lef, rig := tree.Split(i)
			lef.Traverse(lessThan(i))
			rig.Traverse(atLeast(i))
		}
	}
}
