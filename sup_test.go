package splaytree

import "testing"
import "math/rand"
import "time"

func TestSupremum(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []int {
		nums := rand.Perm(upper)
		items := make([]int, upper)
		for i := 0; i < upper; i++ {
			items[i] = nums[i]
		}
		return items
	}

	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if _, ok := tree.Supremum(-1); ok {
		t.Errorf("supremum -1")
	}

	for i := 1; i < 12; i++ {
		tree = NewSplayTree(func(l, r int) bool {
			return l < r
		})
		tree.InsertAll(perm(i))
		for k := -1; k <= i; k++ {
			sup, ok := tree.Supremum(k)
			if !ok && i > 0 && k+1 < i {
				t.Errorf("supremum %v %v nil", i, k)
			}
			if ok && sup != k+1 {
				t.Errorf("supremum %v %v = %v", i, k, sup)
			}
		}
	}
}

func TestInfimum(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []int {
		nums := rand.Perm(upper)
		items := make([]int, upper)
		for i := 0; i < upper; i++ {
			items[i] = nums[i]
		}
		return items
	}

	tree := NewSplayTree(func(l, r int) bool {
		return l < r
	})
	if _, ok := tree.Infimum(-1); ok {
		t.Errorf("infimum -1")
	}

	for i := 1; i < 12; i++ {
		tree = NewSplayTree(func(l, r int) bool {
			return l < r
		})
		tree.InsertAll(perm(i))
		for k := -1; k <= i; k++ {
			inf, ok := tree.Infimum(k)
			if !ok && k > 0 {
				t.Errorf("infimum %v %v nil", i, k)
			}
			if ok && inf != k-1 {
				t.Errorf("infimum %v %v = %v", i, k, inf)
			}
		}
	}
}
